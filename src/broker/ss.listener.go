package broker

import (
	"encoding/json"
	"exex-chart/src/config"
	"exex-chart/src/context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func ssConnectToServer() (*websocket.Conn, error) {
	config := config.GetConfig()

	url := url.URL{
		Scheme: config.Broker.Ss.Scheme,
		Path:   config.Broker.Ss.Path,
		Host:   config.Broker.Ss.Host,
	}
	log.Printf("Сonnecting to SS ws %s", url.String())

	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		return nil, err
	}

	subMessage := SsSubMessage{
		Method: "subscribe",
		Ch:     "trades",
		Params: SsParams{
			Symbols: config.Symbols,
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

func ssListenAndServe(c *websocket.Conn) {
	for {
		_, msg, err := c.ReadMessage()

		if err != nil {
			log.Printf("Error read SS ws: %v", err)
			c.Close()
			return
		}

		var message SsMsgTrades
		err = json.Unmarshal(msg, &message)

		if err != nil {
			fmt.Printf("Error parse message SS ws: %v\n", err)
		} else {

			if message.Update != nil {
				for symbol, trades := range *message.Update {
					for _, trade := range trades {
						msg := context.TradeChanel{
							Timestamp: trade.T,
							Id:        trade.I,
							Price:     trade.P,
							Quantity:  trade.Q,
							Side:      trade.S,
							Symbol:    symbol,
						}

						context.BroadcastTrade <- &msg
					}
				}
			}
		}
	}
}

func SSListener() {
	const maxAttempts = 5
	attempts := 0

	for {
		c, err := ssConnectToServer()
		if err != nil {
			log.Printf("Error connect SS ws: %v", err)
		} else {
			ssListenAndServe(c)
			attempts = 0
		}

		if attempts++; attempts >= maxAttempts {
			fmt.Fprintf(os.Stderr, "Maximum number of connection SS ws attempts exceeded: %v\n", err)
			os.Exit(1)
			break
		}

		log.Printf("Attempting to reconnect after 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}
