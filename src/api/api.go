package api

import (
	"exex-chart/src/_core/context"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func InitApi() {
	port := ":" + context.Config.Api.Port
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc(routes.API.V1.GET_CHART, getChartHendler)

	log.Println("STARTING REST API SERVER ON " + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running"))
}
