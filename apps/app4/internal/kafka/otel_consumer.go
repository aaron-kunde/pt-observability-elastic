package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func startSpan(ctx context.Context, msg *kafka.Message) (context.Context, trace.Span) {
	parentSpanContext := otel.GetTextMapPropagator().Extract(ctx, NewMessageCarrier(msg))
	newCtx, span := tracer.Start(parentSpanContext, fmt.Sprintf("%v receive", msg.Topic))
	addTraceContext(newCtx, msg)

	return newCtx, span
}
