package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Params struct {
	Symbols []string `json:"symbols"`
	Limit   *int     `json:"limit"`
}

type SubMessage struct {
	Method  string  `json:"method"`
	Ch      *string `json:"ch,omitempty"`
	Channel *string `json:"channel,omitempty"`
	Params  *Params `json:"params,omitempty"`
	Options *Params `json:"options,omitempty"`
}

type HandlerParams struct {
	Method string `json:"method"`
	Ch     string `json:"ch"`
	Params Params `json:"params"`
}

type Client struct {
	ID    string
	Conn  *websocket.Conn
	Mutex sync.Mutex
}

type TradeOrder struct {
	T int64  `json:"t"`
	I int64  `json:"i"`
	P string `json:"p"`
	Q string `json:"q"`
	S string `json:"s"`
}

type TradeOrderList map[string][]TradeOrder

type TradeMessage struct {
	Ch       string          `json:"ch"`
	Update   *TradeOrderList `json:"update,omitempty"`
	Snapshot *TradeOrderList `json:"snapshot,omitempty"`
}

type RsiOrderList map[string]float64

type RsiMessage struct {
	Channel string       `json:"channel"`
	Data    RsiOrderList `json:"data"`
}
