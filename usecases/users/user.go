package users_case

import (
	kafka_producer "project2/repository/kafka"
	"project2/usecases"
	"time"

	"github.com/jmoiron/sqlx"
)

func New(c Configuration, d Depencency) usecases.UserUseCase {
	return &usecase{c, d}
}

type Configuration struct {
	Timeout time.Duration
}

type Depencency struct {
	Postgresql    *sqlx.DB
	KafkaProducer kafka_producer.Contract
}

type usecase struct {
	Configuration
	Depencency
}
