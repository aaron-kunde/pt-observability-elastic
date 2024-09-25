package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
)

func addTraceContext(ctx context.Context, msg *kafka.Message) {
	otel.GetTextMapPropagator().Inject(ctx, NewMessageCarrier(msg))
}
