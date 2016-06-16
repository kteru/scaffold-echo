package handler

import (
	"net/http"

	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
)

// Logout handler
func Logout(c echo.Context) error {
	ctx := core.GetContext(c)

	delete(ctx.Session.Values, core.SessionKeyUserID)
	ctx.Session.Options.MaxAge = -1

	ctx.SaveSession(c)
	return c.String(http.StatusOK, "")
}
