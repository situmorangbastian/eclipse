package eclipse

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

func Timeout(duration time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, cancel := context.WithTimeout(c.Request().Context(), duration)
			defer cancel()

			newRequest := c.Request().WithContext(ctx)
			c.SetRequest(newRequest)

			return next(c)
		}
	}
}
