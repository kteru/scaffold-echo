package api

import (
	"github.com/kteru/scaffold-echo/api/handler"
	"github.com/kteru/scaffold-echo/api/middleware"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

// route configures routes
func (s *Server) route(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(echoMw.Recover())
	e.Use(middleware.Session())

	g := e.Group("/api/v1")
	{
		authMws := []echo.MiddlewareFunc{}
		g.POST("/login", handler.Login, authMws...)
		g.POST("/logout", handler.Logout, authMws...)

		normalMws := []echo.MiddlewareFunc{
			middleware.AuthUser(),
		}
		g.GET("/hello", handler.Hello, normalMws...)
	}
}
