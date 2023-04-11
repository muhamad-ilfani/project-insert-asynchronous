package app

import (
	"context"
	"project2/delivery"
	kafka_subs "project2/delivery/kafka"
	kafka_producer "project2/repository/kafka"
	"project2/usecases"
	users_case "project2/usecases/users"
	"time"
)

func (x *App) initService(ctx context.Context) (err error) {
	timeout := 55 * time.Second

	kafka_producer_repo := kafka_producer.NewKafkaProducer(x.kafkaProducer)

	usercase := users_case.New(
		users_case.Configuration{
			Timeout: timeout,
		},
		users_case.Depencency{
			Postgresql:    x.DB,
			KafkaProducer: kafka_producer_repo,
		},
	)

	kafka_subs.InitSubscriptions(ctx, kafka_subs.Usecase{UserCase: usercase})

	delivery.NewEchoHandler(ctx, x.Echo, struct {
		usecases.UserUseCase
	}{
		usercase,
	})

	return nil
}
