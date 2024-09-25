package main

import (
	"context"
	"net/http"
	"pt.observability.elastic/app4/internal/kafka"
	log "pt.observability.elastic/app4/internal/logging"
	"pt.observability.elastic/app4/internal/metrics"
	"pt.observability.elastic/app4/internal/rest"
	"pt.observability.elastic/app4/internal/traces"
	"strconv"
)

var port = 8084

func main() {
	ctx := context.Background()
	log.Info(ctx, "Starting App using Go ")
	traces.Init(ctx)
	metrics.Init(ctx)

	rest.RegisterApiHandler()

	kafka.Listen(ctx)
	http.ListenAndServe(":"+strconv.Itoa(port), traces.NewHTTPHandler())
	log.Info(ctx, "Listening on port: ", port)
}
