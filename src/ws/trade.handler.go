package ws

import (
	"encoding/json"
	"exex-chart/src/_core/context"
	"sync"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var tradeExists = struct{}{}
var tradeSubscribersMutex sync.Mutex
var tradeSubscribers = make(map[string]map[*Client]struct{})

var historyTrade = make(map[string][]context.TradeChanel)

func TradeHandler(client *Client, message HandlerParams) {
	switch message.Method {
	case "subscribe":
		tradeSendSnapshot(client, message.Params.Symbols, message.Params.Limit)
		tradeSubscribe(client, message)
	case "unsubscribe":
		tradeUnsubscribe(client, message)
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
			saveTrade(msg)
			clients := tradeSubscribers[msg.Symbol]

			for client := range clients {
				client.Mutex.Lock()
				client.Conn.WriteMessage(websocket.TextMessage, jsonData)
				client.Mutex.Unlock()
			}
		}
	}
}

func saveTrade(msg *context.TradeChanel) {
	_, isExist := historyTrade[msg.Symbol]

	if isExist == false {
		historyTrade[msg.Symbol] = []context.TradeChanel{*msg}
		return
	}

	historyTrade[msg.Symbol] = append(historyTrade[msg.Symbol], *msg)
	if len(historyTrade[msg.Symbol]) > context.Config.Storage.Trade_history_limit {
		historyTrade[msg.Symbol] = historyTrade[msg.Symbol][1:]
	}
}

func getHistory(symbol string, limit int) []context.TradeChanel {
	history, isExist := historyTrade[symbol]
	slice := len(history) - limit

	if isExist == false {
		return []context.TradeChanel{}
	}

	if slice < 0 {
		slice = 0
	}

	return history[slice:]
}

func tradeSendSnapshot(client *Client, symbols []string, limit *int) {
	var l int
	var snapshot TradeOrderList
	snapshot = make(TradeOrderList)

	if limit == nil {
		l = 500
	} else {
		l = *limit
	}

	for _, symbol := range symbols {
		history := getHistory(symbol, l)
		tradeOrders := make([]TradeOrder, len(history))

		for i, trade := range history {
			tradeOrders[i] = TradeOrder{
				T: trade.Timestamp,
				I: trade.Id,
				P: trade.Price,
				Q: trade.Quantity,
				S: trade.Side,
			}
		}

		snapshot[symbol] = tradeOrders
	}

	message := TradeMessage{
		Ch:       "trades",
		Snapshot: &snapshot,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Error("Error serializing message trade chanel", err)
		return
	}

	client.Mutex.Lock()
	defer client.Mutex.Unlock()
	client.Conn.WriteMessage(websocket.TextMessage, jsonData)
}

func tradeSubscribe(client *Client, message HandlerParams) {
	tradeSubscribersMutex.Lock()
	defer tradeSubscribersMutex.Unlock()

	for _, symbol := range message.Params.Symbols {
		if _, ok := tradeSubscribers[symbol]; !ok {
			tradeSubscribers[symbol] = make(map[*Client]struct{})
		}
		tradeSubscribers[symbol][client] = tradeExists
	}
}

func tradeUnsubscribe(client *Client, message HandlerParams) {
	tradeSubscribersMutex.Lock()
	defer tradeSubscribersMutex.Unlock()

	for _, symbol := range message.Params.Symbols {
		for s := range tradeSubscribers {
			if symbol == s {
				delete(tradeSubscribers[symbol], client)
			}
		}
	}
}

func TradeDropSubscriber(client *Client) {
	tradeSubscribersMutex.Lock()
	defer tradeSubscribersMutex.Unlock()

	for symbol := range tradeSubscribers {
		delete(tradeSubscribers[symbol], client)
	}
}
