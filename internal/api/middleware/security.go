package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/helmet"
)

func SecurityMiddleware() fiber.Handler {
	return helmet.New()
}
