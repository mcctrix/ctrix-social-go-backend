package controllers

import "github.com/gofiber/fiber/v3"

func GetUser() fiber.Handler {
	return func(c fiber.Ctx) error {

		return c.JSON(map[string]string{
			"test": "reached to /user",
		})
	}
}
