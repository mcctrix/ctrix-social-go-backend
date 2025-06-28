package bookmarks

import (
	"github.com/gofiber/fiber/v3"
	repo "github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database/repositories"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
)

func GetBookmark() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Locals("userID").(string)
		limit := utils.QueryLimit(c.Query("limit", "5"))

		bookmarks, err := repo.GetBookmark(userID, limit)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(bookmarks)
	}
}
func CreateBookmark() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Locals("userID").(string)
		postID := c.Params("postID")

		err := repo.CreateBookmark(userID, postID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "bookmark created successfully",
		})
	}
}
func DeleteBookmark() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Locals("userID").(string)
		postID := c.Params("postID")

		err := repo.DeleteBookmark(userID, postID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "bookmark deleted successfully",
		})
	}
}