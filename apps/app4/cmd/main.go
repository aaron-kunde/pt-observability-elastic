package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main()  {
	log.Info("Starting app4")

	http.HandleFunc("/api-1", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("Calling API 1");

		log.Infof("Requeseted API-1:%s", request.URL)
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

	http.ListenAndServe(":8084", nil)
}