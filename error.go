package eclipse

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func Error() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			if e, ok := err.(*echo.HTTPError); ok {
				if e.Code >= http.StatusInternalServerError {
					log.Error(e.Message)
				}

				return echo.NewHTTPError(e.Code, strings.ToLower(e.Message.(string)))
			}

			// Check error based on error type
			switch errors.Cause(err).(type) {
			case ConstraintError:
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			case NotFoundError:
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			case ConflictError:
				return echo.NewHTTPError(http.StatusConflict, err.Error())
			default:
				log.Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
			}
		}
	}
}
