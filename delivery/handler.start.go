package delivery

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
)

func welcome(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": WelcomeMsg,
		})
	}
}
