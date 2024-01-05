package storage

import (
	"exex-chart/src/_core/context"
	"exex-chart/src/_core/pg"
	"fmt"
	"strings"
)

func PatchCandles(db *pg.DB, symbol string, period string, candles []context.CandleCanel) error {

	tableName, err := CandleTable(symbol, period)

	if err != nil {
		return err
	}

	var valueStrings []string
	var valueArgs []interface{}

	for i, candle := range candles {
		valueIdx := i * 5
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", valueIdx+1, valueIdx+2, valueIdx+3, valueIdx+4, valueIdx+5))
		valueArgs = append(valueArgs, candle.StartTime, candle.Open, candle.Low, candle.High, candle.Close)
	}

	sql := fmt.Sprintf(`
		INSERT INTO "%s" AS c ("ts", "open", "low", "high", "close")
		VALUES %s
		ON CONFLICT ("ts")
		DO UPDATE SET
		"low" = LEAST(EXCLUDED.low, c.low),
		"high" = GREATEST(EXCLUDED.high, c.high),
		"close" = EXCLUDED.close
	`, tableName, strings.Join(valueStrings, ","))

	_, err = db.Exec(pg.Ctx, sql, valueArgs...)
	if err != nil {
		return err
	}

	return nil
}
