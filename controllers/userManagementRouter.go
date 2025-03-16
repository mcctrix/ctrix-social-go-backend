package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/db"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
)

func GetCurrentUserProfile() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println(err)
			return c.Status(401).SendString("unable to fetch user!")
		}
		db, err := db.DBConnection()
		if err != nil {
			fmt.Println(err)
			return c.Status(500).SendString("unable to fetch user!")
		}

		return c.SendString("Get current user profile")
	}
}
