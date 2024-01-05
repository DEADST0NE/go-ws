package storage

import "time"

type ResGetChandleCacheKey struct {
	hKey   string
	rowKey *string
}

type ParamsFindCandles struct {
	Symbol string
	Period string
	From   time.Time
	To     time.Time
}
