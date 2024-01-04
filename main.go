package main

import (
	"exex-chart/src/_core/context"
	"exex-chart/src/_core/redis"
	"exex-chart/src/broker"
	"exex-chart/src/candles"
	"exex-chart/src/rsi"
	"exex-chart/src/ws"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("START PROJECT")

	// CONTEXT
	context.Init()

	// REDIS
	redis.Init()

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
