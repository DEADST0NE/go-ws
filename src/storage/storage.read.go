package storage

import (
	"encoding/json"
	"exex-chart/src/_core/context"
	"exex-chart/src/_core/pg"
	"fmt"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"
)

func FindCandles(params ParamsFindCandles) ([]CandleOrder, error) {
	tenMinutesAgo := time.Now().Add(-10 * time.Minute)
	tableName, err := CandleTable(params.Symbol, params.Period)

	if err != nil {
		return []CandleOrder{}, err
	}

	db, err := pg.Connect()

	if err != nil {
		return []CandleOrder{}, err
	}

	defer db.Close()

	var candles []CandleOrder

	if params.To.After(tenMinutesAgo) {
		cache, err := GetCandleCache(params.Symbol, params.Period)

		if err == nil {
			for _, data := range *cache {
				candleStr := data

				if candleStr == "" {
					continue
				}

				var candle context.CandleCanel
				err := json.Unmarshal([]byte(candleStr), &candle)
				if err != nil {
					log.Errorf("ERROR DESERIALIZING STRING IN CANDLECANEL: %v\n", err)
					continue
				}

				if params.To.After(candle.StartTime) && candle.StartTime.Before(params.From) {
					data := CandleOrder{
						Time:  candle.StartTime.Unix(),
						Open:  candle.Open,
						Close: candle.Close,
						High:  candle.High,
						Low:   candle.Low,
					}

					candles = append(candles, data)
				}
			}
		} else {
			log.Errorf("ERROR GET CANDLE FROM REDIS: %v\n", err)
			return []CandleOrder{}, err
		}
	}

	sql := fmt.Sprintf(`
		SELECT ts, open, close, high, low
		FROM %s
		WHERE ts BETWEEN $1 AND $2
		ORDER BY ts ASC
	`, tableName)

	rows, err := db.Query(pg.Ctx, sql, params.From, params.To)

	if err != nil {
		log.Errorf("ERROR GET CANDLE FROM DB: %v\n", err)
		return nil, err
	}

	for rows.Next() {
		var candle CandleOrder
		var time time.Time
		err := rows.Scan(&time, &candle.Open, &candle.Close, &candle.High, &candle.Low)

		if err != nil {
			log.Errorf("ERROR GET CANDLE FROM DB: %v", err)
			continue
		}
		candle.Time = time.Unix()
		candles = append(candles, candle)
	}
	if err := rows.Err(); err != nil {
		log.Errorf("ERROR READING ROWS: %v", err)
	}

	sort.Slice(candles, func(i, j int) bool {
		return candles[i].Time < candles[j].Time
	})

	return candles, nil
}

type CandleOrder struct {
	Time  int64   `json:"time"`
	Open  float64 `json:"open"`
	Close float64 `json:"close"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
}
