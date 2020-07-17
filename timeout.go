package eclipse

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

// ContextTimeout returns custom middleware for Echo that set maximum HTTP response time
// before considered timeout with duration d.
func ContextTimeout(d time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, cancel := context.WithTimeout(c.Request().Context(), d)
			defer cancel()

			newRequest := c.Request().WithContext(ctx)
			c.SetRequest(newRequest)

			return next(c)
		}
	}
}
