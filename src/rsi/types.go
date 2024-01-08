package rsi

type Rsi struct {
	Values     []float64
	Usmma      *float64
	Dsmma      *float64
	Filled     bool
	Symbol     string
	Period     string
	Interval   int
	FillBlanks bool
	rsi        *float64
}
