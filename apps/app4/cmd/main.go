package main

import (
	"net/http"
	log "pt.observability.elastic/app4/internal/logging"
	"pt.observability.elastic/app4/internal/rest"
	"strconv"
)

var port = 8084

func main()  {
	log.Info("Starting app4. Listening on port: ", port)
	rest.RegisterApiHandler()

	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}