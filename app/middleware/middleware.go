package middleware

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func LogIncomingAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hostname, _ := os.Hostname()
		log.Println("Incoming request from", c.Request().RemoteAddr)
		log.Println("Hostname", hostname)

		return next(c)
	}
}
