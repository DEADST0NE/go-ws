package rsi

func RsiCreate(symbol string, period string, interval int, fillBlanks bool) Rsi {
	rsi := Rsi{
		Symbol:     symbol,
		Period:     period,
		Interval:   interval,
		Filled:     false,
		FillBlanks: fillBlanks,
	}

	return rsi
}

func rsiPush(rsi *Rsi, value float64) *float64 {
	rsi.Values = append(rsi.Values, value)
	if rsi.Filled {
		allRsi[rsi.Period][rsi.Symbol] = *rsi
		return rsiRecalculate(rsi)
	}
	if len(rsi.Values) > rsi.Interval {
		rsiRecalculate(rsi)
		allRsi[rsi.Period][rsi.Symbol] = *rsi
		return nil
	}
	allRsi[rsi.Period][rsi.Symbol] = *rsi
	return nil
}

func rsiReplacelast(rsi *Rsi, value float64) *float64 {
	if len(rsi.Values) > 0 {
		rsi.Values[len(rsi.Values)-1] = value
	} else {
		rsi.Values = append(rsi.Values, value)
	}
	if rsi.Filled {
		return rsiRecalculate(rsi)
	}
	return nil
}

func rsiRecalculate(rsi *Rsi) *float64 {
	if !rsi.Filled {
		if len(rsi.Values) < rsi.Interval+1 {
			return nil
		}

		rangeValues := rsi.Values[:rsi.Interval]
		var up []float64
		var down []float64
		up = append(up, 0)
		down = append(down, 0)

		for i := 1; i < rsi.Interval; i++ {
			diff := rangeValues[i] - rangeValues[i-1]
			if diff > 0 {
				up = append(up, diff)
				down = append(down, 0)
			} else {
				up = append(up, 0)
				down = append(down, -diff)
			}
		}

		uSMMAVal := sum(up) / float64(rsi.Interval)
		dSMMAVal := sum(down) / float64(rsi.Interval)

		rsi.Usmma = &uSMMAVal
		rsi.Dsmma = &dSMMAVal
		rsi.Filled = true
		rsi.Values = rsi.Values[rsi.Interval-1:]
	}

	if rsi.Usmma == nil || rsi.Dsmma == nil || *rsi.Usmma == 0 || *rsi.Dsmma == 0 {
		return nil
	}

	if len(rsi.Values) < 2 {
		return nil
	}

	uSMMA := *rsi.Usmma
	dSMMA := *rsi.Dsmma

	for i := 1; i < len(rsi.Values); i++ {
		diff := rsi.Values[i] - rsi.Values[i-1]

		var up float64
		var down float64
		if diff > 0 {
			up = diff
			down = 0
		} else {
			down = -diff
			up = 0
		}

		uSMMA = (up + uSMMA*(float64(rsi.Interval)-1)) / float64(rsi.Interval)
		dSMMA = (down + dSMMA*(float64(rsi.Interval)-1)) / float64(rsi.Interval)
	}

	rs := uSMMA / dSMMA
	rsi.Usmma = &uSMMA
	rsi.Dsmma = &dSMMA

	rsi.Values = rsi.Values[len(rsi.Values)-2:]
	res := 100 - 100/(1+rs)
	rsi.rsi = &res

	return &res
}

func sum(slice []float64) float64 {
	total := 0.0
	for _, value := range slice {
		total += value
	}
	return total
}
