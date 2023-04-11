package users_case

import (
	"context"
	"project2/repository"
	us "project2/repository/users"
	"project2/usecases"
	"time"

	"github.com/rs/zerolog/log"
)

func (x *usecase) RegisterUser(
	ctx context.Context, req usecases.RegisterUserData) (
	res usecases.RegisterUserResponse, httpcode int, err error,
) {
	var timeX time.Time

	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() {
			err = new(repository.SQLTransaction).EndTx(tx, err)
		}()
	}

	defer func() {
		if err != nil {
			//log.Err(err).Msg("Error Register")
			kafkaMessage := repository.NotifyRegistrationRequest{
				ID:        req.ID,
				Customer:  req.Customer,
				Quantity:  req.Quantity,
				Price:     req.Price,
				TimeStamp: timeX,
			}
			if err = x.KafkaProducer.NotifyRegistration(ctx, kafkaMessage); err != nil {
				log.Err(err).Msg("Error Notify")
			}
		}
	}()

	layoutFormat := "2006-01-02 15:04:05"

	timeX, _ = time.Parse(layoutFormat, req.TimeStamp)

	userPG := us.NewRepository(tx)

	request := repository.RegisterUserRequest{
		ID:        req.ID,
		Customer:  req.Customer,
		Quantity:  req.Quantity,
		Price:     req.Price,
		TimeStamp: timeX,
	}

	_, httpcode, err = userPG.RegisterUser(ctx, request)
	if err != nil {

		return res, httpcode, err
	}

	return res, httpcode, err
}
