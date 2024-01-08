package ws

import (
	"encoding/json"
	"exex-chart/src/_core/context"
	"exex-chart/src/_core/redis"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var rsiExists = struct{}{}
var rsiSubscribersMutex sync.Mutex
var rsiSubscribers = make(map[string]map[*Client]struct{})
var lastRsi = make(map[string]float64)

func RsiHandler(client *Client, message HandlerParams) {
	switch message.Method {
	case "subscribe":
		rsiSubscribe(client, message)
	case "unsubscribe":
		rsiUnsubscribe(client, message)
	default:
		SendMessage(client.Conn, "Method not implemented")
	}
}

func GetCacheLastRsi(period string, symbol string) float64 {
	key := getKey(period, symbol)
	msg, err := redis.Client.Get(redis.Ctx, key).Result()

	if err == redis.Nil {
		return 0
	} else if err != nil {
		log.Error("ERROR: GET RSI CACHE")
		return 0
	}

	rsi, err := strconv.ParseFloat(msg, 64)
	if err != nil {
		log.Error("ERROR: CONVERTING RSI TO FLOAT64: ", err)
		return 0
	}

	return rsi
}

func CacheRsi(msg *context.RsiCanel) {
	key := getKey(msg.Period, msg.Symbol)
	redis.Client.Set(redis.Ctx, key, msg.Rsi, 0)
}

func SendRsiUpdate() {
	for {
		msg := <-context.BroadcastRsiWS
		CacheRsi(msg)
		var update RsiOrderList
		update = make(RsiOrderList)
		update[msg.Symbol] = msg.Rsi

		message := RsiMessage{
			Channel: "rsi/" + msg.Period,
			Data:    update,
		}

		key := getKey(msg.Period, msg.Symbol)

		lastRsi[key] = msg.Rsi

		jsonData, err := json.Marshal(message)
		if err != nil {
			log.Error("ERROR: SERIALIZING MESSAGE RSI CHANEL", err)
			continue
		}

		if msg != nil {
			clients := rsiSubscribers[key]

			for client := range clients {
				client.Conn.WriteMessage(websocket.TextMessage, jsonData)
			}
		}
	}
}

func getKey(period string, symbol string) string {
	return "RSI:" + symbol + ":" + period
}

func rsiSubscribe(client *Client, message HandlerParams) {
	rsiSubscribersMutex.Lock()
	defer rsiSubscribersMutex.Unlock()

	parts := strings.Split(message.Ch, "/")
	var period string
	if len(parts) > 1 {
		period = parts[1]
	} else {
		period = context.Config.Ws.Default.Rsi
	}

	for _, symbol := range message.Params.Symbols {
		key := getKey(period, symbol)
		if _, ok := rsiSubscribers[key]; !ok {
			rsiSubscribers[key] = make(map[*Client]struct{})
		}
		rsiSubscribers[key][client] = rsiExists
	}
	rsiSendSnapshot(client, period, message.Params.Symbols)
}

func rsiSendSnapshot(client *Client, period string, symbols []string) {
	var update RsiOrderList
	update = make(RsiOrderList)

	for _, symbol := range symbols {
		key := getKey(period, symbol)

		rsi, isExist := lastRsi[key]
		if isExist {
			update[symbol] = rsi
		} else {
			rsi = GetCacheLastRsi(period, symbol)
			if rsi > 0 {
				update[symbol] = rsi
			}
		}
	}

	message := RsiMessage{
		Channel: "rsi/" + period,
		Data:    update,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Error("ERROR SERIALIZING MESSAGE RSI CHANEL", err)
		return
	}
	client.Conn.WriteMessage(websocket.TextMessage, jsonData)
}

func rsiUnsubscribe(client *Client, message HandlerParams) {
	rsiSubscribersMutex.Lock()
	defer rsiSubscribersMutex.Unlock()

	parts := strings.Split(message.Ch, "/")
	var period string
	if len(parts) > 1 {
		period = parts[1]
	} else {
		period = context.Config.Ws.Default.Rsi
	}

	for _, symbol := range message.Params.Symbols {
		for s := range rsiSubscribers {
			if symbol == s {
				key := getKey(period, symbol)
				delete(rsiSubscribers[key], client)
			}
		}
	}
}

func RsiDropSubscriber(client *Client) {
	rsiSubscribersMutex.Lock()
	defer rsiSubscribersMutex.Unlock()

	for key := range rsiSubscribers {
		delete(rsiSubscribers[key], client)
	}
}
