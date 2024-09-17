package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	log "pt.observability.elastic/app4/internal/logging"
	"pt.observability.elastic/app4/internal/rest"
	"strconv"
)

var port = 8084

func main() {
	log.Info("Starting App using Go. Listening on port: ", port)

	http.Handle("/actuator/prometheus", promhttp.Handler())

	rest.RegisterApiHandler()

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
