package broker

import (
	"encoding/json"
	"exex-chart/src/config"
	"exex-chart/src/context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func coreConnectToServer() (*websocket.Conn, error) {
	config := config.GetConfig()

	url := url.URL{
		Scheme: config.Broker.Core.Scheme,
		Path:   config.Broker.Core.Path,
		Host:   config.Broker.Core.Host,
	}
	log.Printf("Сonnecting to CORE ws %s", url.String())

	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		return nil, err
	}

	subMessage := CoreSubMessage{
		Method: "subscribeTrades",
		Options: CoreOptions{
			Env: config.Broker.Core.Env,
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
		log.Printf("Error parsing timestamp: %v", err)
		return nil, err
	}

	timestamp := t.Unix()

	id, err := strconv.ParseInt(data.Id, 10, 64)

	if err != nil {
		log.Printf("Error parsing ID: %v", err)
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
			log.Printf("Error read CORE ws: %v", err)
			c.Close()
			return
		}

		var message CoreMsgTrades
		err = json.Unmarshal(msg, &message)

		if err != nil {
			fmt.Printf("Error parse message CORE ws: %v\n", err)
		} else {
			if message.requestId == "" {
				continue
			}

			msg, err := coreParseTrade(message)

			if err != nil {
				fmt.Printf("Error parse CoreTrades to TradeChanel ws: %v\n", err)
				continue
			}

			context.BroadcastTrade <- msg
		}
	}
}

func CoreListener() {
	const maxAttempts = 5
	attempts := 0

	for {
		c, err := coreConnectToServer()
		if err != nil {
			log.Printf("Error connect CORE ws: %v", err)
		} else {
			coreListenAndServe(c)
			attempts = 0
		}

		if attempts++; attempts >= maxAttempts {
			fmt.Fprintf(os.Stderr, "Maximum number of connection CORE ws attempts exceeded: %v\n", err)
			os.Exit(1)
			break
		}

		log.Printf("Attempting to reconnect after 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}
