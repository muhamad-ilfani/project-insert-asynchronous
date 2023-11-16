package app

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (ox *App) initKafka() error {
	var err error

	brokerAddr := fmt.Sprintf("%s:%s", os.Getenv("BROKER_HOST"), os.Getenv("BROKER_PORT"))

	conf := &kafka.ConfigMap{
		"bootstrap.servers": brokerAddr,
	}

	/*ox.kafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    brokerAddr,
		"group.id":             "project2",
		"max.poll.interval.ms": 30 * 1000 * 60, // 30 minutes
	})*/

	if err != nil {
		return err
	}

	ox.kafkaProducer, err = kafka.NewProducer(conf)
	if err != nil {
		return err
	}

	return nil
}
