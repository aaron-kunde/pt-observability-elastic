package traces

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"net/http"
	"os"
	log "pt.observability.elastic/app4/internal/logging"
)

func Init(ctx context.Context) {
	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(newProvider(newExporter(ctx)))
}

func NewHTTPHandler() http.Handler {
	return otelhttp.NewHandler(http.DefaultServeMux, "/")

}

func newExporter(ctx context.Context) sdktrace.SpanExporter {
	var endpoint = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")

	if endpoint == "" {
		endpoint = "localhost:8200"
	}
	// Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
	exporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint(endpoint),
	)

	if err != nil {
		log.Error(nil, err)
	}
	return exporter
}

func newProvider(exporter sdktrace.SpanExporter) *sdktrace.TracerProvider {
	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
		),
	)

	if err != nil {
		log.Error(nil, err)
	}
	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(r),
	)
}
