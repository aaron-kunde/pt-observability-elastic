package main

import (
	"context"
	"net/http"
	log "pt.observability.elastic/app4/internal/logging"
	"pt.observability.elastic/app4/internal/metrics"
	"pt.observability.elastic/app4/internal/rest"
	"pt.observability.elastic/app4/internal/traces"
	"strconv"
)

var port = 8084

func main() {
	log.Info("Starting App using Go. Listening on port: ", port)
	traces.SetupOTelSDK(context.Background())
	metrics.Init()

	rest.RegisterApiHandler()

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
