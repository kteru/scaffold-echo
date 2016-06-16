package middleware

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
)

// RequestID generates the request ID
func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			attr := core.GetAttribute(c)

			b := make([]byte, 12)
			rand.Read(b)
			requestID := base64.URLEncoding.EncodeToString(b)

			attr.RequestID = requestID
			c.Response().Header().Set("X-Request-Id", requestID)

			return next(c)
		}
	}
}
