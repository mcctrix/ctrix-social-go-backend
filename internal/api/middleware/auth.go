package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/auth"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		auth_token := c.Cookies("auth_token")
		if auth_token == "" {
			return c.Status(400).JSON(map[string]string{
				"error": "User not found!",
			})
		}

		jwtToken, err := auth.GetJwtToken(auth_token)
		if err != nil {
			fmt.Println(err)
			return fiber.ErrInternalServerError
		}

		if !jwtToken.Valid {
			return c.Status(400).JSON(map[string]string{
				"error": "Invalid Token",
			})
		}
		c.Locals("userID", auth.GetClaimData(jwtToken, "aud"))

		err = c.Next()
		if err != nil {
			return err
		}
		return nil
	}
}