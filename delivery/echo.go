package delivery

import (
	"context"
	"project2/usecases"

	"github.com/labstack/echo"
)

const (
	TokenIsRequired    = "Token must be provided"
	SuccessMsg         = "Success"
	WelcomeMsg         = "welcome"
	FailedToUnmarshall = "Failed to Unmarshall"
	FailedToRegister   = "Failed to Register"
)

type echoObject struct {
	*echo.Echo
	UseCase
}

type UseCase struct {
	usecases.UserUseCase
}

func NewEchoHandler(ctx context.Context, c *echo.Echo, uc UseCase) {
	obj := &echoObject{c, uc}
	obj.initRoute(ctx)

	obj.Logger.Fatal(obj.Start(":8000"))
}
