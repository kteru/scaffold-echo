package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kteru/scaffold-echo/api/core"
	"github.com/kteru/scaffold-echo/api/request"
	"github.com/labstack/echo"
)

// Login handler
func Login(c echo.Context) error {
	cfg := core.GetConfig(c)
	attr := core.GetAttribute(c)

	reqLogin := request.Login{}
	if err := json.NewDecoder(c.Request().Body()).Decode(&reqLogin); err != nil {
		cfg.Logger.Printf("[%s] [%s] %v", core.LogRoll, attr.RequestID, err)
		return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	if reqLogin.Username == nil {
		cfg.Logger.Printf("[%s] [%s] `username` not found", core.LogRoll, attr.RequestID)
		return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	if reqLogin.Password == nil {
		cfg.Logger.Printf("[%s] [%s] `password` not found", core.LogRoll, attr.RequestID)
		return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	user, err := cfg.DBHandler.AuthUser(*reqLogin.Username, *reqLogin.Password)
	if err != nil {
		cfg.Logger.Printf("[%s] [%s] %v", core.LogRoll, attr.RequestID, err)
		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	if user == nil {
		cfg.Logger.Printf("[%s] [%s] username or password is incorrect", core.LogRoll, attr.RequestID)
		return c.String(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	cfg.Logger.Printf("[%s] [%s] authorized: \"%s\"", core.LogRoll, attr.RequestID, user.Username)
	attr.Session.Values[core.SessionKeyUserID] = user.ID
	attr.SaveSession(c)
	return c.String(http.StatusOK, "")
}
