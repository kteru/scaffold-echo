package handler

import (
	"fmt"
	"net/http"

	"github.com/kteru/scaffold-echo/api/core"
	"github.com/kteru/scaffold-echo/api/response"
	"github.com/labstack/echo"
)

// Hello ...
func Hello(c echo.Context) error {
	ctx := core.GetContext(c)

	resHello := response.Hello{
		Message: fmt.Sprintf("Hello, %s!", ctx.User.Username),
	}

	ctx.SaveSession(c)
	return c.JSON(http.StatusOK, resHello)
}
