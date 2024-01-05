package storage

import (
	"encoding/json"
	"exex-chart/src/_core/context"
	"exex-chart/src/_core/redis"
	"time"

	log "github.com/sirupsen/logrus"
)

func InitWatchCandleSave() {
	log.Info("INIT WATCH SAVE CANDLE")
	for {
		msg := <-context.BroadcastCandleSave
		SetCandleCache(msg)
	}
}

func SetCandleCache(candle *context.CandleCanel) {
	key := GetChandleCacheKey(candle.Symbol, candle.Period, &candle.StartTime)

	candleJSON, errParse := json.Marshal(candle)
	if errParse != nil {
		log.Errorf("ERROR SERIALIZING CANDLE: %v", errParse)
		return
	}

	candleJSONStr := string(candleJSON)

	err := redis.Client.HSet(redis.Ctx, key.hKey, *key.rowKey, candleJSONStr).Err()

	if err != nil {
		log.Errorf("ERROR SAVE CANDLE CACHE: %v", err)
		return
	}
}

func GetCandleCache(symbol string, period string) (*map[string]string, error) {
	key := GetChandleCacheKey(symbol, period, nil)
	candels, err := redis.Client.HGetAll(redis.Ctx, key.hKey).Result()

	if err != nil {
		log.Errorf("ERROR GET CANDLE CACHE: %v", err)
		return nil, err
	}

	return &candels, nil
}

func DeleteCandleCache(candle *context.CandleCanel) error {
	key := GetChandleCacheKey(candle.Symbol, candle.Period, &candle.StartTime)
	_, err := redis.Client.HDel(redis.Ctx, key.hKey, *key.rowKey).Result()

	if err != nil {
		log.Errorf("ERROR DELETE CANDLE CACHE: %v", err)
		return err
	}

	return nil
}

func GetChandleCacheKey(symbol string, period string, date *time.Time) ResGetChandleCacheKey {
	var res ResGetChandleCacheKey
	hashKey := "candle:" + symbol + ":" + period
	res.hKey = hashKey

	if date != nil {
		ts := date.Format(time.RFC3339)
		res.rowKey = &ts
	}

	return res
}
