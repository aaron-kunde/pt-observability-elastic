package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"os"
	log "pt.observability.elastic/app4/internal/logging"
	"pt.observability.elastic/app4/internal/metrics"
	"time"
)

type Producer struct {
	conn         *kafka.Conn
	topicCounter metrics.Counter
}

var (
	tracer          = otel.Tracer("")
	topicOutName    = "topic4"
	applicationName = initApplicationName()
	kafkaProducer   = Producer{
		conn: initConnection(),
		topicCounter: metrics.NewCounter(
			"topic_out_counter",
			map[string]string{"it_1": "it-2"}),
	}
)

func initApplicationName() string {
	var applicationName = os.Getenv("SERVICE_NAME")

	if applicationName == "" {
		applicationName = "app4"
	}
	return applicationName
}

func initConnection() *kafka.Conn {
	partition := 0
	address := os.Getenv("KAFKA_BOOTSTRAP-SERVERS")

	if address == "" {
		address = "localhost:9092"
	}
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topicOutName, partition)

	if err != nil {
		log.Error("Failed to dial leader:", err)
	}
	return conn
}

func Send(ctx context.Context, apiName string, data uint64) {
	_, span := tracer.Start(ctx, "KafkaProducer#Send")
	defer span.End()

	log.Info(fmt.Sprintf("Send data to topic %s: %d", topicOutName, data))

	err := kafkaProducer.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		log.Error("Failed to set deadline:", err)
	}
	_, err = kafkaProducer.conn.WriteMessages(
		kafka.Message{Value: []byte(fmt.Sprintf("%s;%s;data:%d", applicationName, apiName, data))},
	)

	if err != nil {
		log.Error("Failed to write messages:", err)
	}
	kafkaProducer.topicCounter.Increment()
}
