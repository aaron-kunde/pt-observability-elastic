package kafka

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/segmentio/kafka-go"
	"os"
	log "pt.observability.elastic/app4/internal/logging"
	"time"
)

type Producer struct {
	conn         *kafka.Conn
	topicCounter prometheus.Counter
}

var (
	topicOutName    = "topic4"
	applicationName = os.Getenv("SERVICE_NAME")
	kafkaProducer   = Producer{
		conn: initConnection(),
		topicCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name:        fmt.Sprintf("%s_topic_out_counter", applicationName),
			Namespace:   "app4",
			ConstLabels: prometheus.Labels{"it_1": "it-2"},
		}),
	}
)

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

func Send(apiName string, data uint64) {
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
	kafkaProducer.topicCounter.Inc()
}
