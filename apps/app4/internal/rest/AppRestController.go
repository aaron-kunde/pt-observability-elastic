package rest

import (
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"net/http"
	"pt.observability.elastic/app4/internal/db"
	"pt.observability.elastic/app4/internal/kafka"
	log "pt.observability.elastic/app4/internal/logging"
	"pt.observability.elastic/app4/internal/metrics"
)

var (
	api1Counter = metrics.NewCounter("api_1_counter", map[string]string{"it_1": "it-2"})
	api2Counter = metrics.NewCounter("api_2_counter", map[string]string{"it_1": "it-2"})
	api3Counter = metrics.NewCounter("api_3_counter", map[string]string{"it_1": "it-2"})
	tracer      = otel.Tracer("")
)

func RegisterApiHandler() {
	http.HandleFunc("/api-1", func(writer http.ResponseWriter, request *http.Request) {
		ctx, span := tracer.Start(request.Context(), "AppRestController#HandleFunc /api-1")
		defer span.End()

		var apiName = "API 1"
		log.Info(fmt.Sprintf("Calling %s", apiName))
		api1Counter.Increment(ctx)

		var count = api1Counter.Count()
		kafka.Send(ctx, apiName, count)

		dataEntity := db.DataEntity{Data: fmt.Sprintf("AppRestController-1: %d", count)}
		db.Save(ctx, dataEntity)

		// Write response
		fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
	})

	http.HandleFunc("/api-2", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("Calling API 2")
		api2Counter.Increment(request.Context())

		// Must be called, before writing a diffferent response
		err := errors.New("An unexpected error occurred")
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Error(err)
	})

	http.HandleFunc("/api-3", func(writer http.ResponseWriter, request *http.Request) {
		ctx, span := tracer.Start(request.Context(), "AppRestController#HandleFunc /api-3")
		defer span.End()

		log.Info("Calling API 3")
		api3Counter.Increment(request.Context())

		dataEntity := db.DataEntity{Data: fmt.Sprintf("AppRestController-3: %d", api3Counter.Count())}
		db.Save(ctx, dataEntity)

		// Write response
		fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
	})
}
