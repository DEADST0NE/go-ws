package storage

import (
	"fmt"
	"strings"
)

func parseSymbol(symbol string) (SymbolDescriptor, error) {
	symbol = strings.ToUpper(symbol)
	isPerp := false

	if strings.HasSuffix(symbol, "_PERP") {
		isPerp = true
		symbol = strings.ReplaceAll(symbol, "_PERP", "")
	}

	var base string
	if strings.HasSuffix(symbol, "USDT") {
		base = "usdt"
		symbol = strings.ReplaceAll(symbol, "USDT", "")
	} else {
		return SymbolDescriptor{}, fmt.Errorf("unsupported trading symbol: %s", symbol)
	}

	return SymbolDescriptor{
		Quote:  symbol,
		Base:   base,
		IsPerp: isPerp,
	}, nil
}

func CandleTable(symbol, period string) (string, error) {
	desc, err := parseSymbol(symbol)
	if err != nil {
		return "", err
	}

	perp := ""
	if desc.IsPerp {
		perp = "_perp"
	}

	return fmt.Sprintf("candle_%s_%s%s_%s",
		strings.ToLower(desc.Quote),
		strings.ToLower(desc.Base),
		perp,
		strings.ToLower(period),
	), nil
}

type SymbolDescriptor struct {
	Quote  string
	Base   string
	IsPerp bool
}
