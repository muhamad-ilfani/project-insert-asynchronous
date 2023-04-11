package app

import (
	"context"
	"fmt"
	"os"
	"project2/repository/users/query"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func (x *App) initDB(ctx context.Context) (err error) {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	sqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)

	x.DB, err = sqlx.Connect("postgres", sqlconn)
	if err != nil {
		return err
	}

	x.DB.SetConnMaxLifetime(time.Second * 14400)

	return x.DB.PingContext(ctx)
}

func (x *App) initSchema(ctx context.Context) (err error) {
	_, err = x.DB.Exec(query.CreateSchema)
	if err != nil {
		return err
	}

	return nil
}

func (x *App) initTable(ctx context.Context) (err error) {
	_, err = x.DB.Exec(query.CreateTable)
	if err != nil {
		return err
	}

	return nil
}
