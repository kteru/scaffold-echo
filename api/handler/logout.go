package handler

import (
	"net/http"

	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
)

// Logout handler
func Logout(c echo.Context) error {
	attr := core.GetAttribute(c)

	delete(attr.Session.Values, core.SessionKeyUserID)
	attr.Session.Options.MaxAge = -1

	attr.SaveSession(c)
	return c.String(http.StatusOK, "")
}
