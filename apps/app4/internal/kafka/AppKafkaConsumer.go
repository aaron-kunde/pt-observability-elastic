package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"net/http"
	"os"
	"pt.observability.elastic/app4/internal/db"
	log "pt.observability.elastic/app4/internal/logging"
	"pt.observability.elastic/app4/internal/metrics"
)

var in_topics = []string{"topic1", "topic2", "topic3"}

var topicCounter = metrics.NewCounter(
	"topic_in_counter",
	map[string]string{"it_1": "it-2"})

func Listen(ctx context.Context) {
	for _, topic := range in_topics {
		go listenOnTopic(ctx, topic)
	}
}

func listenOnTopic(ctx context.Context, topic string) {
	address := os.Getenv("KAFKA_BOOTSTRAP-SERVERS")

	if address == "" {
		address = "localhost:9092"
	}

	restOutUrl := os.Getenv("REST_OUT_URL")

	if restOutUrl == "" {
		restOutUrl = "http://localhost:8081/api-3"
	}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{address},
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	log.Info(ctx, "Listen on topic: ", topic)

	for {
		message, err := reader.FetchMessage(ctx)

		if err != nil {
			log.Error(ctx, err)
			break
		}
		ctx, span := startSpan(ctx, &message)
		log.Info(ctx, fmt.Sprintf("Fetch data from topic %s: %s=%s", topic, message.Key, message.Value))
		topicCounter.Increment(ctx)

		save(ctx, fmt.Sprintf("AppKafkaConsumer: %d", topicCounter.Count()))

		response, _ := http.Get(restOutUrl)
		log.Info(ctx, fmt.Sprintf("Call REST URL: %s, result: %s", restOutUrl, response))
		span.End()
	}

	if err := reader.Close(); err != nil {
		log.Error(ctx, "failed to close reader:", err)
	}
	log.Info(ctx, "close reader: "+topic)
}

func save(ctx context.Context, data string) {
	dataEntity := db.DataEntity{Data: data}
	log.Info(ctx, fmt.Sprintf("Write data to database: %s", dataEntity))
	db.Save(ctx, dataEntity)
}
