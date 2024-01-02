package rest

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RegisterApiHandler() {
	http.HandleFunc("/api-1", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("Calling API 1");

		// Write response
		fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
	})

	http.HandleFunc("/api-2", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("Calling API 2");
		// Must be called, before writing the response
		http.Error(writer, "An unexpected error occurred", http.StatusInternalServerError)

		// Write response
		fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
	})

	http.HandleFunc("/api-3", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("Calling API 3");

		// Write response
		fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
	})

}

