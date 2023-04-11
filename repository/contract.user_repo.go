package repository

import (
	"context"
	"time"
)

type UserRepo interface {
	RegisterUser(
		ctx context.Context, req RegisterUserRequest) (
		res RegisterUserResponse, httpcode int, err error,
	)
}

type RegisterUserRequest struct {
	ID        int64
	Customer  string
	Quantity  int64
	Price     float64
	TimeStamp time.Time
}

type RegisterUserResponse struct {
}
