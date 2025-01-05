package producer

import (
	"context"
	"encoding/json"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
	"time"
)

type KProducer struct {
	Writer *kafka.Writer
}

func NewKProducer() *KProducer {
	return &KProducer{
		Writer: &kafka.Writer{
			Addr:     kafka.TCP(strings.Split(config.GetString("KAFKA_BROKERS", "localhost:9092"), ",")...),
			Balancer: &kafka.LeastBytes{},
		},
	}
}

// Produce is a function that sends a message to the Kafka broker.
func (k *KProducer) Produce(key, topic string, data interface{}) {
	msgBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error when marshal message: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Retry 3 times before giving up. This is to handle transient errors.
	for retries := 0; retries < 3; retries++ {
		log.Printf("Attempting to produce message: topic=%s key=%s value=%s (attempt %d)",
			topic, key, string(msgBytes), retries+1)

		err = k.Writer.WriteMessages(ctx, kafka.Message{
			Topic: topic,
			Key:   []byte(key),
			Value: msgBytes,
		})

		if err == nil {
			log.Println("Message produced successfully")
			return
		}

		log.Printf("Error producing message: %v", err)
		time.Sleep(2 * time.Second) // Wait before retrying
	}
}
