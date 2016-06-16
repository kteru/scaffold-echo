package middleware

import (
	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
)

// SetContext sets a empty Context to echo.Context
func SetContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &core.Context{}

			c.Set("ctx", ctx)
			return next(c)
		}
	}
}
