package producer

import (
	"context"
	"encoding/json"
	"github.com/ngdangkietswe/swe-auth-service/configs"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type KProducer struct {
	Writer *kafka.Writer
}

func NewKProducer(topic string) *KProducer {
	return &KProducer{
		Writer: &kafka.Writer{
			Addr:     kafka.TCP(configs.GlobalConfig.KafkaBrokers),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

// Produce is a function that sends a message to the Kafka broker.
func (k *KProducer) Produce(key string, data interface{}) {
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
			k.Writer.Topic, key, string(msgBytes), retries+1)

		err = k.Writer.WriteMessages(ctx, kafka.Message{
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