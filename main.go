package main

import (
	"exex-chart/src/broker"
	"exex-chart/src/candles"
	"exex-chart/src/context"
	"exex-chart/src/rsi"
	"exex-chart/src/ws"
	"fmt"
)

func main() {
	fmt.Println("START PROJECT")
	// CONTEXT
	context.Init()

	// DATASOURCE
	go broker.SSListener()
	go broker.CoreListener()

	// CANDLES
	go candles.InitCron()
	go candles.InitTradeChanal()

	// Rsi
	go rsi.InitNewCandleChanal()

	// WS SERVER
	go ws.SendTradeUpdate()
	go ws.SendRsiUpdate()
	ws.WsServer()
}
