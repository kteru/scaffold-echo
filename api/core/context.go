package core

import (
	"github.com/gorilla/sessions"
	"github.com/kteru/scaffold-echo/db/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// Context represents contexts of the current HTTP request
type Context struct {
	RequestID string
	Session   *sessions.Session
	User      *model.User
}

// GetContext gets the Context from echo.Context
func GetContext(c echo.Context) *Context {
	return c.Get("ctx").(*Context)
}

// SaveSession saves the session
func (ctx *Context) SaveSession(c echo.Context) {
	r := c.Request().(*standard.Request).Request
	w := c.Response().(*standard.Response).ResponseWriter
	ctx.Session.Save(r, w)
}
