package handler

import (
	"fmt"
	"net/http"

	"github.com/kteru/scaffold-echo/api/core"
	"github.com/kteru/scaffold-echo/api/response"
	"github.com/labstack/echo"
)

// Hello handler
func Hello(c echo.Context) error {
	attr := core.GetAttribute(c)

	resHello := response.Hello{
		Message: fmt.Sprintf("Hello, %s!", attr.User.Username),
	}

	attr.SaveSession(c)
	return c.JSON(http.StatusOK, resHello)
}
