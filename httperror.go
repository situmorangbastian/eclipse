package eclipse

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

// ErrMiddleware returns custom middleware for Echo that generate HTTP error response
// with HTTP status code.
func ErrMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			if e, ok := err.(*echo.HTTPError); ok {
				switch e.Code {
				case http.StatusNotFound:
				default:
					log.Error(e.Message)
				}

				return echo.NewHTTPError(e.Code, e.Message)
			}

			err = errors.Cause(err)

			// Check error based on error type
			switch err.(type) {
			case ConstraintError:
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			case NotFoundError:
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			case ConflictError:
				return echo.NewHTTPError(http.StatusConflict, err.Error())
			default:
				log.Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
			}
		}
	}
}
