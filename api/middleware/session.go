package middleware

import (
	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// Session ...
func Session() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cfg := core.GetConfig(c)
			ctx := core.GetContext(c)

			r := c.Request().(*standard.Request).Request

			ctx.Session, _ = cfg.SessionStore.Get(r, core.SessionName)
			ctx.Session.Options.MaxAge = core.SessionExpire

			return next(c)
		}
	}
}
