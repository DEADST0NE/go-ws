package cron

import (
	"encoding/json"
	"exex-chart/src/_core/context"
	"exex-chart/src/storage"
	"time"

	"exex-chart/src/_core/pg"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func InitCandleCron() {
	log.Info("INIT CANDLE CRON")

	c := cron.New()

	_, err := c.AddFunc("*/10 * * * *", func() {
		cronCandleJob()
	})

	if err != nil {
		log.Errorf("ERROR INIT JOB CANDLE CRON: %v", err)
	}

	c.Start()

	select {}
}

func cronCandleJob() {
	log.Info("START CRON TRANSFER CANDLE TO DB :" + time.Now().String())

	db, err := pg.Connect()

	if err != nil {
		return
	}

	defer db.Close()

	var candles []context.CandleCanel

	for _, pereod := range context.Config.Candle.Periods {
		for _, symbol := range context.Config.Symbols {
			cache, err := storage.GetCandleCache(symbol, pereod)

			if err != nil {
				continue
			}

			if cache != nil {
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

					candles = append(candles, candle)
				}

				if len(candles) > 0 {
					err := storage.PatchCandles(db, symbol, pereod, candles)

					if err != nil {
						log.Errorf("ERROR INSER CANDLES: %v\n", err)
						continue
					}

					for _, candle := range candles {
						storage.DeleteCandleCache(&candle)
					}

					candles = []context.CandleCanel{}
				} else {
					log.Warn("NOT INFO CANDLES IN CACHE")
				}
			}
		}
	}
}
