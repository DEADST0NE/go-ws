package ws

import (
	"encoding/json"
	"exex-chart/src/context"
	"sync"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var exists = struct{}{}
var subscribersMutex sync.Mutex
var subscribers = make(map[string]map[*Client]struct{})

func TradeHandler(client *Client, message SubMessage) {
	switch message.Method {
	case "subscribe":
		subscribe(client, message)
	case "unsubscribe":
		unsubscribe(client, message)
	default:
		SendMessage(client.Conn, "Method not implemented")
	}
}

func SendTradeUpdate() {
	for {
		msg := <-context.BroadcastTradeWS

		var update TradeOrderList
		update = make(TradeOrderList)
		update[msg.Symbol] = []TradeOrder{
			{T: msg.Timestamp, I: msg.Id, P: msg.Price, Q: msg.Quantity, S: msg.Side},
		}

		message := TradeMessage{
			Ch:     "trades",
			Update: &update,
		}

		jsonData, err := json.Marshal(message)
		if err != nil {
			log.Error("Error serializing message trade chanel", err)
			continue
		}

		if msg != nil {
			clients := subscribers[msg.Symbol]

			for client := range clients {
				client.Conn.WriteMessage(websocket.TextMessage, jsonData)
			}
		}
	}
}

func subscribe(client *Client, message SubMessage) {
	subscribersMutex.Lock()
	defer subscribersMutex.Unlock()

	for _, symbol := range message.Params.Symbols {
		if _, ok := subscribers[symbol]; !ok {
			subscribers[symbol] = make(map[*Client]struct{})
		}
		subscribers[symbol][client] = exists
	}
}

func unsubscribe(client *Client, message SubMessage) {
	subscribersMutex.Lock()
	defer subscribersMutex.Unlock()

	for _, symbol := range message.Params.Symbols {
		for s := range subscribers {
			if symbol == s {
				delete(subscribers[symbol], client)
			}
		}
	}
}

func TradeDropSubscriber(client *Client) {
	subscribersMutex.Lock()
	defer subscribersMutex.Unlock()

	for symbol := range subscribers {
		delete(subscribers[symbol], client)
	}
}
