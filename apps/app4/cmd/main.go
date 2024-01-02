package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"pt.observability.elastic/app4/internal/rest"
	"strconv"
)

func main()  {
	var port = 8084
	log.Info("Starting app4. Listening on port: ", port)
	rest.RegisterApiHandler()

	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}