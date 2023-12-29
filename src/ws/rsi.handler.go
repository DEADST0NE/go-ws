package ws

import (
	"encoding/json"
	"exex-chart/src/context"
	"sync"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var rsiExists = struct{}{}
var rsiSubscribersMutex sync.Mutex
var rsiSubscribers = make(map[string]map[*Client]struct{})

func RsiHandler(client *Client, message SubMessage) {
	switch message.Method {
	case "subscribe":
		rsiSubscribe(client, message)
	case "unsubscribe":
		rsiUnsubscribe(client, message)
	default:
		SendMessage(client.Conn, "Method not implemented")
	}
}

func SendRsiUpdate() {
	for {
		msg := <-context.BroadcastRsiWS

		var update RsiOrderList
		update = make(RsiOrderList)
		update[msg.Symbol] = msg.Rsi

		message := RsiMessage{
			Channel: "rsi" + msg.Period,
			Data:    update,
		}

		jsonData, err := json.Marshal(message)
		if err != nil {
			log.Error("Error serializing message rsi chanel", err)
			continue
		}

		if msg != nil {
			clients := rsiSubscribers[msg.Symbol]

			for client := range clients {
				client.Conn.WriteMessage(websocket.TextMessage, jsonData)
			}
		}
	}
}

func rsiSubscribe(client *Client, message SubMessage) {
	rsiSubscribersMutex.Lock()
	defer rsiSubscribersMutex.Unlock()

	for _, symbol := range message.Params.Symbols {
		if _, ok := rsiSubscribers[symbol]; !ok {
			rsiSubscribers[symbol] = make(map[*Client]struct{})
		}
		rsiSubscribers[symbol][client] = rsiExists
	}
}

func rsiUnsubscribe(client *Client, message SubMessage) {
	rsiSubscribersMutex.Lock()
	defer rsiSubscribersMutex.Unlock()

	for _, symbol := range message.Params.Symbols {
		for s := range rsiSubscribers {
			if symbol == s {
				delete(rsiSubscribers[symbol], client)
			}
		}
	}
}

func RsiDropSubscriber(client *Client) {
	rsiSubscribersMutex.Lock()
	defer rsiSubscribersMutex.Unlock()

	for symbol := range rsiSubscribers {
		delete(rsiSubscribers[symbol], client)
	}
}
