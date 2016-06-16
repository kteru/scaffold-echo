package middleware

import (
	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// Session prepares to use the session
func Session() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cfg := core.GetConfig(c)
			attr := core.GetAttribute(c)

			r := c.Request().(*standard.Request).Request

			attr.Session, _ = cfg.SessionStore.Get(r, core.SessionName)
			attr.Session.Options.MaxAge = core.SessionExpire

			return next(c)
		}
	}
}
