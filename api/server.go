package api

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/gorilla/sessions"
	"github.com/kteru/scaffold-echo/api/core"
	"github.com/kteru/scaffold-echo/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// ServerConfig ...
type ServerConfig struct {
	ListenHost string
	ListenPort int
	Logger     *log.Logger
	Secret     []byte
	DBHandler  db.Handler
}

// Server ...
type Server struct {
	ServerConfig
	config *core.Config
}

// NewServer ...
func NewServer(c *ServerConfig) (*Server, error) {
	if c.Logger == nil {
		return nil, fmt.Errorf("`Logger` not found")
	}

	if len(c.Secret) == 0 {
		b := make([]byte, 64)
		rand.Read(b)
		c.Secret = b
	}

	cfg := &core.Config{
		Logger:       c.Logger,
		SessionStore: sessions.NewCookieStore(c.Secret),
		DBHandler:    c.DBHandler,
	}
	if err := core.ValidConfig(cfg); err != nil {
		return nil, err
	}

	s := Server{
		ServerConfig: *c,
		config:       cfg,
	}

	return &s, nil
}

// Serve ...
func (s *Server) Serve() error {
	e := echo.New()
	s.route(e)

	std := standard.New(fmt.Sprintf("%s:%d", s.ListenHost, s.ListenPort))
	std.SetHandler(e)

	return std.Start()
}
