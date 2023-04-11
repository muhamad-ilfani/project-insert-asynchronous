package main

import (
	"context"
	"project2/app"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app.Run(context.Background())
}
