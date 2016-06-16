package middleware

import (
	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
)

// SetAttribute sets a empty Attribute to echo.Context
func SetAttribute() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			attr := &core.Attribute{}

			c.Set("attr", attr)
			return next(c)
		}
	}
}
