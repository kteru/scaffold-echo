package core

import (
	"github.com/gorilla/sessions"
	"github.com/kteru/scaffold-echo/db/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// Attribute represents attributes of the current HTTP request
type Attribute struct {
	RequestID string
	Session   *sessions.Session
	User      *model.User
}

// GetAttribute gets the Attribute from the Context
func GetAttribute(c echo.Context) *Attribute {
	return c.Get("attr").(*Attribute)
}

// SaveSession saves the session
func (attr *Attribute) SaveSession(c echo.Context) {
	r := c.Request().(*standard.Request).Request
	w := c.Response().(*standard.Response).ResponseWriter
	attr.Session.Save(r, w)
}
