package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	db "github.com/mcctrix/ctrix-social-go-backend/db/v1"
)

func GetFeed() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Locals("userID").(string)

		posts, err := db.GetPostFeed(userID)
		if err != nil {
			fmt.Println("error while fetching feed: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch feed!")
		}
		return c.JSON(posts)
	}
}
func GetFollowRecommendation() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Locals("userID").(string)

		recommendation, err := db.GetFollowRecommendation(userID)
		if err != nil {
			fmt.Println("error while fetching follow recommendation: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch follow recommendation!")
		}
		return c.JSON(recommendation)
	}
}
