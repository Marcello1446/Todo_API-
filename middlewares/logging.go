package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
)

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		method := c.Request().Method
		path := c.Request().URL.Path
		status := c.Response().Status
		duration := time.Since(start)

		fmt.Printf("[%s] %s %d %s\n", method, path, status, duration)
		return next(c)
	}
}
