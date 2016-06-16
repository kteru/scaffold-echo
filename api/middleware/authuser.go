package middleware

import (
	"net/http"

	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
)

// AuthUser authenticates user
func AuthUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cfg := core.GetConfig(c)
			ctx := core.GetContext(c)

			userID, ok := ctx.Session.Values[core.SessionKeyUserID]
			if !ok {
				return c.String(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			}

			user, err := cfg.DBHandler.FindUserByID(userID.(uint))
			if err != nil {
				return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
			if user == nil {
				return c.String(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			}

			ctx.User = user

			return next(c)
		}
	}
}
