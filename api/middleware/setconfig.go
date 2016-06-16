package middleware

import (
	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
)

// SetConfig sets the core.Config to echo.Context
func SetConfig(cfg *core.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("cfg", cfg)
			return next(c)
		}
	}
}
