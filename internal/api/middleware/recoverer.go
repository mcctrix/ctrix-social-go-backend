package middleware

import (
	"github.com/gofiber/fiber/v3"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
)

// RecovererMiddleware recovers from panics and returns a 500 status code.
func RecovererMiddleware() fiber.Handler {
	return recoverer.New()
}
