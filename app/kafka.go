package app

import "github.com/confluentinc/confluent-kafka-go/kafka"

func (ox *App) initKafka() error {
	var err error

	conf := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	ox.kafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    "localhost:9092",
		"group.id":             "project2",
		"max.poll.interval.ms": 30 * 1000 * 60, // 30 minutes
	})

	if err != nil {
		return err
	}

	ox.kafkaProducer, err = kafka.NewProducer(conf)
	if err != nil {
		return err
	}

	return nil
}
