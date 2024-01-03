package rest

import (
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	log "pt.observability.elastic/app4/internal/logging"
)

var (
	namespace = "app4"
	api1Counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "api_1_counter",
		Namespace: namespace,
		ConstLabels: prometheus.Labels{"it_1": "it-2"},
	})
	api2Counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "api_2_counter",
		Namespace: namespace,
		ConstLabels: prometheus.Labels{"it_1": "it-2"},
	})
	api3Counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "api_3_counter",
		Namespace: namespace,
		ConstLabels: prometheus.Labels{"it_1": "it-2"},
	})
)
func RegisterApiHandler() {
	http.HandleFunc("/api-1", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("Calling API 1");
		api1Counter.Inc()

		// Write response
		fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
	})

	http.HandleFunc("/api-2", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("Calling API 2");
		api2Counter.Inc()

		// Must be called, before writing a diffferent response
		err := errors.New("An unexpected error occurred")
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Error(err)
	})

	http.HandleFunc("/api-3", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("Calling API 3");
		api3Counter.Inc()

		// Write response
		fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
	})

}

