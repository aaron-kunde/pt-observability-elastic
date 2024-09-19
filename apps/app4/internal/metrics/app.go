package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"os"
)

type Counter interface {
	Increment(ctx context.Context)
	Count() uint64
}

type counter struct {
	localCounter uint64
	promCounter  prometheus.Counter
	otelCounter  metric.Int64Counter
}

func NewCounter(name string, labels map[string]string) Counter {
	namespace := applicationName()
	int64Counter, _ := otel.Meter("").Int64Counter(namespace + "_" + name)

	return &counter{
		localCounter: 0,
		promCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name:        name,
			Namespace:   namespace,
			ConstLabels: labels,
		}),
		otelCounter: int64Counter,
	}
}

func Init(ctx context.Context) {
	setupPrometheusEndpoint()
	setupOTelSDK(ctx)
}

func applicationName() string {
	var applicationName = os.Getenv("SERVICE_NAME")

	if applicationName == "" {
		applicationName = "app4"
	}
	return applicationName
}

func (counter *counter) Increment(ctx context.Context) {
	counter.promCounter.Inc()
	counter.localCounter++
	counter.otelCounter.Add(ctx, 1)
}

func (counter *counter) Count() uint64 {
	return counter.localCounter
}
