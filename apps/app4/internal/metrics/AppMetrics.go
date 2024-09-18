package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"os"
)

type Counter interface {
	Increment()
	Count() uint64
}

type counter struct {
	cnt         uint64
	promCounter prometheus.Counter
}

func (counter *counter) Increment() {
	counter.promCounter.Inc()
	counter.cnt++
}

func (counter *counter) Count() uint64 {
	return counter.cnt
}

func NewCounter(name string, labels map[string]string) Counter {

	return &counter{
		cnt: 0,
		promCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name:        name,
			Namespace:   initApplicationName(),
			ConstLabels: labels,
		}),
	}
}

func initApplicationName() string {
	var applicationName = os.Getenv("SERVICE_NAME")

	if applicationName == "" {
		applicationName = "app4"
	}
	return applicationName
}
