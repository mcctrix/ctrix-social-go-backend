package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func Login() fiber.Handler {
	return func(c fiber.Ctx) error {
		fmt.Println("Reach in Login Controller!", c.OriginalURL())
		// g.JSON(http.StatusOK, gin.H{
		// 	"test": "working",
		// })
		return nil
	}
}
