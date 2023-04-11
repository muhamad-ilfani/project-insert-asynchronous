package repository

import (
	"context"
	"time"
)

type NotifyRegistration interface {
	NotifyRegistration(
		ctx context.Context, req NotifyRegistrationRequest) (
		err error,
	)
}

type NotifyRegistrationRequest struct {
	ID        int64
	Customer  string
	Quantity  int64
	Price     float64
	TimeStamp time.Time
}

type NotifyRegistrationResponse struct{}
