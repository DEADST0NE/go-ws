package candles

import (
	"exex-chart/src/_core/context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

var allCandles = make(map[string]map[string]context.CandleCanel)
var lastPrice = make(map[string]float64)

func InitCron() {
	log.Info("INIT CANDLE GENERATOR")

	c := cron.New(cron.WithSeconds()) //, cron.WithLogger(cron.VerbosePrintfLogger(log.StandardLogger()))

	for _, period := range context.Config.Candle.Periods {
		period := period
		spec, err := periodToCronSpec(period)

		if err != nil {
			log.Error("ERROR CONFIG CRON JOB:", err)
			os.Exit(1)
		}

		_, err = c.AddFunc(spec, func() {
			cronJob(period, context.Config.Symbols)
		})

		if err != nil {
			log.Error("ERROR INIT JOB: ", period, " ", spec)
		}
	}

	c.Start()

	select {}
}

func InitTradeChanal() {
	for {
		trade := <-context.BroadcastTradeCandle
		processTrade(trade)
	}
}

func cronJob(period string, sumbols []string) {
	currentDate := time.Now()
	log.Info("START CANDLE JOB PERIOD:", period, "TIME:", currentDate.String())

	closedCandle := make(map[string]bool)
	candles := allCandles[period]

	for symbol, candle := range candles {
		d1 := currentDate.Truncate(time.Second)
		d2 := candle.EndTime.Truncate(time.Second)

		if d2.Before(d1) || d2.Equal(d1) {
			fire(&candle)
			closedCandle[symbol] = true
			delete(allCandles[period], symbol)
		} else {
			log.Error("Error candle symbol: ", symbol, " ", period, " ", "endTime <= new ", d2, " ", "<=", " ", d1)
		}
	}

	for _, symbol := range sumbols {
		_, isExist := closedCandle[symbol]

		if isExist == false {
			log.Warn("Not trade for candle: ", symbol, " ", period)

			price, exPrice := lastPrice[symbol]

			if exPrice {
				dates, err := calculateTime(currentDate, period)

				if err != nil {
					log.Error("Error calculateTime:", err)
				}

				candle := context.CandleCanel{
					Symbol:    symbol,
					Open:      price,
					High:      price,
					Low:       price,
					Close:     price,
					Period:    period,
					StartTime: dates.PrevTime,
					EndTime:   dates.EndTime,
				}

				fire(&candle)
			} else {
				log.Warn("Not info last price candle: ", symbol, " ", period)
			}
		}
	}
}

func fire(candle *context.CandleCanel) {
	copyCandle := context.CandleCanel{
		Open:      candle.Open,
		High:      candle.High,
		Low:       candle.Low,
		Close:     candle.Close,
		Symbol:    candle.Symbol,
		StartTime: candle.StartTime,
		EndTime:   candle.EndTime,
		Period:    candle.Period,
	}

	rsiStatus, isExist := context.Config.Candle.Rsi_events[copyCandle.Period]
	if isExist && rsiStatus {
		context.BroadcastCandleRsi <- copyCandle
	}

	context.BroadcastCandleSave <- &copyCandle
}

func calculateTime(date time.Time, period string) (ResCalculateTime, error) {
	var result ResCalculateTime

	date = date.Truncate(time.Second)

	switch period {
	case "S15":
		roundedSeconds := date.Second() - date.Second()%15
		result.StartTime = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), roundedSeconds, 0, date.Location())
		result.EndTime = result.StartTime.Add(15 * time.Second)
		result.PrevTime = result.StartTime.Add(-15 * time.Second)
	case "M1":
		result.StartTime = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 0, 0, date.Location())
		result.EndTime = result.StartTime.Add(time.Minute)
		result.PrevTime = result.StartTime.Add(-time.Minute)
	case "M5":
		roundedMinutes := date.Minute() - date.Minute()%5
		result.StartTime = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), roundedMinutes, 0, 0, date.Location())
		result.EndTime = result.StartTime.Add(5 * time.Minute)
		result.PrevTime = result.StartTime.Add(-5 * time.Minute)
	case "M15":
		roundedMinutes := date.Minute() - date.Minute()%15
		result.StartTime = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), roundedMinutes, 0, 0, date.Location())
		result.EndTime = result.StartTime.Add(15 * time.Minute)
		result.PrevTime = result.StartTime.Add(-15 * time.Minute)
	default:
		return ResCalculateTime{}, fmt.Errorf("Not info period for cron")
	}

	return result, nil
}

func periodToCronSpec(timeFrame string) (string, error) {
	switch timeFrame {
	case "S15":
		return "*/15 * * * * *", nil
	case "M1":
		return "0 * * * * *", nil
	case "M5":
		return "0 */5 * * * *", nil
	case "M15":
		return "0 */15 * * * *", nil
	default:
		return "", fmt.Errorf("Not info timeframe for cron")
	}
}

func updateOldCandle(trade *context.TradeChanel) {
	log.Info("Update candle", trade)
	// context.BroadcastCandleRsiUpdate <- trade
}

func updateCandle(candle context.CandleCanel, price float64) {
	if price > candle.High {
		candle.High = price
	}
	if price < candle.Low {
		candle.Low = price
	}
	candle.Close = price

	allCandles[candle.Period][candle.Symbol] = candle
}

func processTrade(trade *context.TradeChanel) {
	price, err := strconv.ParseFloat(trade.Price, 64)
	if err != nil {
		log.Error("Parce price to float64:", err)
		return
	}

	seconds := trade.Timestamp / 1000
	nanoseconds := (trade.Timestamp % 1000) * 1000000
	tradeTime := time.Unix(seconds, nanoseconds)

	lastPrice[trade.Symbol] = price

	for _, period := range context.Config.Candle.Periods {
		dates, err := calculateTime(tradeTime, period)
		candle, isExCandle := allCandles[period][trade.Symbol]

		if isExCandle && candle.EndTime.Before(tradeTime) {
			fire(&candle)
			delete(allCandles[period], trade.Symbol)
			if err != nil {
				log.Warn("Error calculateTime:", err)
			}
		} else {
			if candle.StartTime.After(tradeTime) {
				updateOldCandle(trade)
				return
			}

			if allCandles[period] == nil {
				allCandles[period] = make(map[string]context.CandleCanel)
			}

			if isExCandle == false {
				allCandles[period][trade.Symbol] = context.CandleCanel{
					Open:      price,
					High:      price,
					Low:       price,
					Close:     price,
					Period:    period,
					Symbol:    trade.Symbol,
					StartTime: dates.StartTime,
					EndTime:   dates.EndTime,
				}
				return
			}

			candle = allCandles[period][trade.Symbol]
			updateCandle(candle, price)
		}

	}
}
