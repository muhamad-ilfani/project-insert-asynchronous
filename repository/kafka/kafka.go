package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"project2/repository"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Contract interface {
	repository.NotifyRegistration
}

// NewKafkaProducer.
func NewKafkaProducer(producer *kafka.Producer) Contract {
	return &kafkaProducer{producer: producer}
}

type kafkaProducer struct {
	producer *kafka.Producer
}

// TODO: adding option parameters.
func (x *kafkaProducer) PublishData(ctx context.Context, topic string, payload interface{}) (err error) {
	msgValue, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return x.publishRawData(ctx, topic, msgValue)
}

func (x *kafkaProducer) publishRawData(ctx context.Context, topic string, payload []byte) error {
	if x.producer == nil {
		return errors.New("kafka producer is nil")
	}

	deliveryChan := make(chan kafka.Event)

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
	}

	if err := x.producer.Produce(msg, deliveryChan); err != nil {
		return err
	}

	e := <-deliveryChan
	m, _ := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}

func (x *kafkaProducer) NotifyRegistration(
	/*req*/ ctx context.Context, request repository.NotifyRegistrationRequest) (
	/*res*/ err error,
) {
	topic := "registration.notification.retry"

	return x.PublishData(ctx, topic, request)
}
