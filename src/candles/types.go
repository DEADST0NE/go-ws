package candles

import "time"

type ResCalculateTime struct {
	EndTime   time.Time
	StartTime time.Time
	PrevTime  time.Time
}
