package core

import (
	"fmt"
	"log"

	"github.com/gorilla/sessions"
	"github.com/kteru/scaffold-echo/db"
	"github.com/labstack/echo"
)

// Config is the configuration to process requests
type Config struct {
	Logger       *log.Logger
	SessionStore *sessions.CookieStore
	DBHandler    db.Handler
}

// GetConfig gets the Config from the Context
func GetConfig(c echo.Context) *Config {
	return c.Get("cfg").(*Config)
}

// ValidConfig validates the Config
func ValidConfig(cfg *Config) error {
	switch {
	case cfg == nil:
		return fmt.Errorf("`Config` not found")
	case cfg.Logger == nil:
		return fmt.Errorf("`Logger` not found")
	case cfg.SessionStore == nil:
		return fmt.Errorf("`SessionStore` not found")
	case cfg.DBHandler == nil:
		return fmt.Errorf("`DBHandler` not found")
	}

	return nil
}
