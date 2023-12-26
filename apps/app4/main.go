package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main()  {
	log.Info("Starting app4")

	http.HandleFunc("/api-1", func(writer http.ResponseWriter, request *http.Request) {
		log.Infof("Requeseted API-1:%s", request.URL)
		// Write response
		fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
	})

	http.ListenAndServe(":8084", nil)
}