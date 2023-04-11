package app

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type App struct {
	DB            *sqlx.DB
	Echo          *echo.Echo
	kafkaConsumer *kafka.Consumer
	kafkaProducer *kafka.Producer
}

func Run(ctx context.Context) {
	app := new(App)

	if err := app.initDB(ctx); err != nil {
		log.Fatal(err)
	}

	if err := app.initSchema(ctx); err != nil {
		log.Fatal(err)
	}

	if err := app.initTable(ctx); err != nil {
		log.Fatal(err)
	}

	if err := app.initKafka(); err != nil {
		log.Fatal(err)
	}

	if err := app.initService(ctx); err != nil {
		log.Fatal(err)
	}
}
