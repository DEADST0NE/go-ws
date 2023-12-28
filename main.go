package main

import (
	"exex-chart/src/broker"
	"exex-chart/src/candles"
	"exex-chart/src/ws"
	"fmt"
)

func main() {
	fmt.Println("START PROJECT")

	// DATASOURCE
	go broker.SSListener()
	go broker.CoreListener()

	// CANDLES
	go candles.InitCron()
	go candles.InitTradeChanal()

	// WS SERVER
	go ws.SendTradeUpdate()
	ws.WsServer()
}
