package rsi

import (
	"exex-chart/src/_core/context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var allRsi = make(map[string]map[string]Rsi)
var allLastCandle = make(map[string]map[string]context.CandleCanel)

func InitNewCandleChanal() {
	for _, generator := range context.Config.Rsi {
		for _, symbol := range context.Config.Symbols {
			if allRsi[generator.Period] == nil {
				allRsi[generator.Period] = make(map[string]Rsi)
			}
			allRsi[generator.Period][symbol] = RsiCreate(symbol, generator.Period, generator.Interval, generator.FillBlanks)
		}
	}

	for {
		candle := <-context.BroadcastCandleRsi
		rsi, isExist := allRsi[candle.Period][candle.Symbol]

		if isExist {
			isSend := addCandle(&rsi, &candle)
			if isSend {
				sendChanelRsi(&rsi)
			}
		}
	}
}

func addCandle(rsi *Rsi, candle *context.CandleCanel) bool {
	c := *candle
	var value *float64

	if rsi == nil {
		log.Error("Not found rsi")
		return false
	}

	if allLastCandle[c.Period] == nil {
		allLastCandle[c.Period] = make(map[string]context.CandleCanel)
	}
	lastCandle, idExist := allLastCandle[c.Period][c.Symbol]

	if idExist {
		if c.StartTime.Equal(lastCandle.StartTime) {
			value = rsiReplacelast(rsi, c.Close)
			allLastCandle[c.Period][c.Symbol] = c

			return value != nil
		}
	}

	if rsi.FillBlanks {
		currentCandle, err := fillBlanks(&c, &lastCandle)

		if err != nil {
			return false
		}

		c = *currentCandle
	}

	value = rsiPush(rsi, c.Close)

	allLastCandle[c.Period][c.Symbol] = c
	return value != nil
}

func sendChanelRsi(rsi *Rsi) {
	value := *rsi.rsi

	data := context.RsiCanel{
		Symbol: rsi.Symbol,
		Period: rsi.Period,
		Rsi:    value,
	}

	log.Info("RSI", data)

	context.BroadcastRsiWS <- &data
}
func fillBlanks(currentCandle, previousCandle *context.CandleCanel) (*context.CandleCanel, error) {
	if previousCandle.Close == 0 {
		return currentCandle, nil
	}

	period, err := periodToSeconds(currentCandle.Period)
	if err != nil {
		return nil, err
	}

	leftT := previousCandle.StartTime.Unix()
	rightT := currentCandle.StartTime.Unix()
	diff := rightT - leftT

	if diff > int64(period)*3/2 { // 1.5 times interval
		ts := previousCandle.StartTime.Add(time.Duration(period) * time.Second)
		dummy := *previousCandle // Shallow copy
		dummy.StartTime = ts
		return &dummy, nil
	}

	return currentCandle, nil
}

func periodToSeconds(period string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	value := re.FindString(period)

	if value == "" {
		return 0, fmt.Errorf("invalid format")
	}

	amount, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid number")
	}

	switch {
	case strings.HasPrefix(period, "S"):
		return amount, nil
	case strings.HasPrefix(period, "M"):
		return amount * 60, nil
	case strings.HasPrefix(period, "H"):
		return amount * 60 * 60, nil
	case strings.HasPrefix(period, "D"):
		return amount * 24 * 60 * 60, nil
	case strings.HasPrefix(period, "W"):
		return amount * 7 * 24 * 60 * 60, nil
	default:
		return 0, fmt.Errorf("unknown period format")
	}
}
