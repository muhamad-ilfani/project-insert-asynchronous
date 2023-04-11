package user_repo

import (
	"project2/repository"

	"github.com/jmoiron/sqlx"
)

type List []interface{}

type PostgreSQLConn struct {
	tc *sqlx.Tx
}

type Repository interface {
	repository.UserRepo
}

func NewRepository(tc *sqlx.Tx) Repository { return &PostgreSQLConn{tc} }
