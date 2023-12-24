package ws

import "github.com/gorilla/websocket"

type Params struct {
	Symbols []string `json:"symbols"`
}

type SubMessage struct {
	Method string `json:"method"`
	Ch     string `json:"ch"`
	Params Params `json:"params"`
}

type Client struct {
	ID   string
	Conn *websocket.Conn
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
