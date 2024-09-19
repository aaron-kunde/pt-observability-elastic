package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func setupPrometheusEndpoint() {
	http.Handle("/actuator/prometheus", promhttp.Handler())
}
