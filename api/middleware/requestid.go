package middleware

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
)

// RequestID ...
func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := core.GetContext(c)

			b := make([]byte, 12)
			rand.Read(b)
			requestID := base64.URLEncoding.EncodeToString(b)

			ctx.RequestID = requestID
			c.Response().Header().Set("X-Request-Id", requestID)

			return next(c)
		}
	}
}
