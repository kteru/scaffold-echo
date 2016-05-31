package middleware

import (
	"time"

	"github.com/kteru/scaffold-echo/api/core"
	"github.com/labstack/echo"
)

// Logger ...
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cfg := core.GetConfig(c)
			ctx := core.GetContext(c)

			req := c.Request()
			res := c.Response()

			ra := req.RemoteAddress()
			if ip := req.Header().Get(echo.HeaderXRealIP); ip != "" {
				ra = ip
			} else if ip := req.Header().Get(echo.HeaderXForwardedFor); ip != "" {
				ra = ip
			}

			cfg.Logger.Printf("[%s] [%s] [%s] Started %s %q\n", core.LogRoll, ctx.RequestID, ra, req.Method(), req.URI())

			t1 := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			t2 := time.Now()

			cfg.Logger.Printf("[%s] [%s] [%s] Returning %03d (%v)\n", core.LogRoll, ctx.RequestID, ra, res.Status(), t2.Sub(t1))

			return nil
		}
	}
}
