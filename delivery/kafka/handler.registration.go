package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"project2/usecases"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func callBackRegistration(ctx context.Context, uc usecases.UserUseCase) handler {
	return func(consumer *kafka.Consumer, msg *kafka.Message) (err error) {
		go func() {
			payload := usecases.RegisterUserData{}

			err = json.Unmarshal(msg.Value, &payload)
			if err != nil {
				fmt.Println(err)

			}

			//fmt.Printf("Subscribe %v", payload)

			uc.RegisterUser(ctx, payload)
		}()

		return nil
	}
}
