package context

// TRADE
var BroadcastTradeWS = make(chan *TradeChanel, 1000)
var BroadcastTradeCandle = make(chan *TradeChanel, 1000)

// CANDLE
// var BroadcastCandleSave = make(chan *CandleCanel, 1000)

var BroadcastCandleRsiUpdate = make(chan *TradeChanel, 1000)
var BroadcastCandleRsi = make(chan *CandleCanel, 1000)
