package context

import (
	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo"
)

type FormContext struct {
	echo.Context
	DB *sqlx.DB
}

// FormContextMiddleware converts the echo.Context to context.FormContext.FormContextMiddleware
// It MUST be registered first in the list of middleware.
func FormContextMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &FormContext{c, db}
			return h(cc)
		}
	}
}
