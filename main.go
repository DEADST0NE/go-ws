package main

import (
	"exex-chart/src/broker"
	"exex-chart/src/ws"
	"fmt"
)

func main() {
	fmt.Println("START PROJECT")
	go broker.SSListener()
	go broker.CoreListener()

	go ws.SendTradeUpdate()
	ws.WsServer()
}
