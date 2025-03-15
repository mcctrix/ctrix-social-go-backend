package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Do something here
		fmt.Println("Reach in AuthMiddleware!")
		err := c.Next()
		if err != nil {
			return err
		}
		return nil
	}
}
