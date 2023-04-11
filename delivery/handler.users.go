package delivery

import (
	"context"
	"net/http"
	"project2/usecases"

	"github.com/labstack/echo"
)

func RegisterUser(ctx context.Context, uc usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		form := usecases.RegisterUserRequest{}

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		go func() {
			for _, req := range form.Data {
				uc.RegisterUser(ctx, req)
			}
		}()

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message":   SuccessMsg,
			"RequestID": form.RequestID,
		})
	}
}
