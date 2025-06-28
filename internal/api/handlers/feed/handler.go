package feed

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	repo "github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database/repositories"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
)

func GetFeed() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Locals("userID").(string)
		limit := utils.QueryLimit(c.Query("limit", "5"))

		posts, err := repo.GetPostFeed(userID, limit)
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
		limit := utils.QueryLimit(c.Query("limit", "5"))

		recommendation, err := repo.GetFollowRecommendation(userID, limit)
		if err != nil {
			fmt.Println("error while fetching follow recommendation: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch follow recommendation!")
		}
		return c.JSON(recommendation)
	}
}