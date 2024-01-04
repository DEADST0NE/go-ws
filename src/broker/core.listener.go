package broker

import (
	"encoding/json"
	"exex-chart/src/_core/context"
	"net/url"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

func coreConnectToServer() (*websocket.Conn, error) {

	url := url.URL{
		Scheme: context.Config.Broker.Core.Scheme,
		Path:   context.Config.Broker.Core.Path,
		Host:   context.Config.Broker.Core.Host,
	}
	log.Info("СONNECTING TO CORE WS URL: ", url.String())

	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		return nil, err
	}

	subMessage := CoreSubMessage{
		Method: "subscribeTrades",
		Options: CoreOptions{
			Env: context.Config.Broker.Core.Env,
		},
	}

	messageBytes, err := json.Marshal(subMessage)

	if err != nil {
		return nil, err
	}

	c.WriteMessage(websocket.TextMessage, messageBytes)

	if err != nil {
		return nil, err
	}

	return c, err
}

func coreParseTrade(data CoreMsgTrades) (*context.TradeChanel, error) {
	// Format definition ISO 8601
	layout := "2006-01-02T15:04:05.000Z"

	t, err := time.Parse(layout, data.Timestamp)

	if err != nil {
		log.Error("Error parsing timestamp:", err)
		return nil, err
	}

	timestamp := t.Unix()

	id, err := strconv.ParseInt(data.Id, 10, 64)

	if err != nil {
		log.Error("Error parsing ID:", err)
		return nil, err
	}

	quantity := strconv.Itoa(int(data.Quantity))
	price := strconv.FormatFloat(data.Price, 'f', 6, 64)

	msg := context.TradeChanel{
		Timestamp: timestamp,
		Id:        id,
		Price:     price,
		Quantity:  quantity,
		Side:      data.Side,
		Symbol:    data.Symbol,
	}

	return &msg, nil
}

func coreListenAndServe(c *websocket.Conn) {
	for {
		_, msg, err := c.ReadMessage()

		if err != nil {
			log.Error("Error read CORE ws:", err)
			c.Close()
			return
		}

		var message CoreMsgTrades
		err = json.Unmarshal(msg, &message)

		if err != nil {
			log.Error("Error parse message CORE ws:", err)
		} else {
			if message.RequestId == "" {
				continue
			}

			msg, err := coreParseTrade(message)

			if err != nil {
				log.Error("Error parse CoreTrades to TradeChanel ws:", err)
				continue
			}

			context.BroadcastTradeWS <- msg
			context.BroadcastTradeCandle <- msg
		}
	}
}

func CoreListener() {
	const maxAttempts = 5
	attempts := 0

	for {
		c, err := coreConnectToServer()
		if err != nil {
			log.Error("Error connect CORE ws:", err)
		} else {
			coreListenAndServe(c)
			attempts = 0
		}

		if attempts++; attempts >= maxAttempts {
			log.Error("Maximum number of connection CORE ws attempts exceeded:", err)
			os.Exit(1)
			break
		}

		log.Info("Attempting to reconnect after 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}
