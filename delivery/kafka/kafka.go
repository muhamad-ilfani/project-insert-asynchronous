package kafka

import (
	"context"
	"fmt"
	"os"
	"project2/usecases"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"
)

type handler func(consumer *kafka.Consumer, msg *kafka.Message) (err error)

type Usecase struct {
	UserCase usecases.UserUseCase
}

type SubcribeParams struct {
	Group      string
	Topic      string
	AutoCommit bool
	Handler    handler
}

func InitSubscriptions(
	ctx context.Context,
	features Usecase,
) (err error) {
	messages := make(chan *kafka.Message, 1)

	subscribers := []SubcribeParams{
		{
			"registration",
			"registration.notification.retry",
			true,
			callBackRegistration(ctx, features.UserCase),
		},
	}

	for _, subcriber := range subscribers {
		go func() error {
			fmt.Println("Consumer Subscribe")
			err = subscribe(
				ctx,
				subcriber.Group,
				subcriber.Topic,
				subcriber.AutoCommit,
				subcriber.Handler,
				messages,
			)

			return err
		}()
		if err != nil { //nolint:wsl
			return err
		}
	}

	return nil
}

func subscribe(ctx context.Context, group, topic string, autoCommit bool, h handler, messages chan<- *kafka.Message) error {
	brokerAddr := fmt.Sprintf("%s:%s", os.Getenv("BROKER_HOST"), os.Getenv("BROKER_PORT"))

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    brokerAddr,
		"group.id":             group,
		"max.poll.interval.ms": 1 * 60 * 1000, //30 * 1000 * 60, // 30 minutes
		"enable.auto.commit":   autoCommit,
	})

	defer func() (err error) {
		if err := consumer.Close(); err != nil {
			log.Err(err).Msg("ERROR consumer close")

			return err
		}

		return nil
	}()

	if err != nil {
		return err
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		return err
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Err(err).Msg("err read message")
		}

		select {
		case <-ctx.Done():
			return err
		case messages <- msg:
			//log.Printf("message fetched and sent to a channel: %v \n", string(msg.Value))
		}

		err = h(consumer, msg)
		if err != nil {
			log.Err(err).
				Interface("msg", msg).
				Msg("kafka - message handler")
		}
	}
}
