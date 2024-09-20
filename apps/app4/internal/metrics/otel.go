package metrics

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"os"
	log "pt.observability.elastic/app4/internal/logging"
	"time"
)

func setupOTelSDK(ctx context.Context) {
	otel.SetMeterProvider(newProvider(newExporter(ctx)))
}

func newExporter(ctx context.Context) sdkmetric.Exporter {
	var endpoint = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")

	if endpoint == "" {
		endpoint = "localhost:8200"
	}

	exporter, err := otlpmetrichttp.New(
		ctx,
		otlpmetrichttp.WithInsecure(),
		otlpmetrichttp.WithEndpoint(endpoint),
	)

	if err != nil {
		log.Error(nil, err)
	}
	return exporter
}

func newProvider(exporter sdkmetric.Exporter) *sdkmetric.MeterProvider {

	return sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			sdkmetric.WithInterval(3*time.Second))),
	)
}
