package usecases

import "context"

type UserUseCase interface {
	RegisterUser(
		ctx context.Context, req RegisterUserData) (
		res RegisterUserResponse, httpcode int, err error,
	)
}

type RegisterUserRequest struct {
	RequestID int64              `json:"request_id"`
	Data      []RegisterUserData `json:"data"`
}

type RegisterUserResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type RegisterUserData struct {
	ID        int64   `json:"id"`
	Customer  string  `json:"customer"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
	TimeStamp string  `json:"timestamp"`
	IsRetry   bool    `json:"is_retry"`
}
