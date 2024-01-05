package api

import (
	"encoding/json"
	"exex-chart/src/storage"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var defaultPeriod = "S15"

const dateTimeFormat = "2006-01-02 15:04:05"

func getChartHendler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var msg GetChartDto
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fromTime, err := time.Parse(dateTimeFormat, msg.From)
	if err != nil {
		http.Error(w, "Invalid 'from' date format. Expected YYYY-MM-DD HH:MM:SS.", http.StatusBadRequest)
		return
	}

	toTime, err := time.Parse(dateTimeFormat, msg.To)
	if err != nil {
		http.Error(w, "Invalid 'to' date format. Expected YYYY-MM-DD HH:MM:SS.", http.StatusBadRequest)
		return
	}

	period := msg.Period

	if period == nil {
		period = &defaultPeriod
	}

	var res CandleOrderList = make(CandleOrderList)

	for _, symbol := range msg.Symbols {
		params := storage.ParamsFindCandles{
			Symbol: symbol,
			Period: *period,
			From:   fromTime,
			To:     toTime,
		}

		candles, err := storage.FindCandles(params)

		if err != nil {
			log.Errorf("Error %v:", err)
			res[symbol] = make([]storage.CandleOrder, 0)
			continue
		}

		if candles != nil && len(*candles) > 0 {
			res[symbol] = *candles
		} else {
			res[symbol] = make([]storage.CandleOrder, 0)
		}
	}

	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

type GetChartDto struct {
	Symbols []string `json:"symbols"`
	From    string   `json:"from"`
	To      string   `json:"to"`
	Period  *string  `json:"period"`
}

type CandleOrderList map[string][]storage.CandleOrder
