package controllers

import (
	"github.com/gofiber/fiber/v3"
	db "github.com/mcctrix/ctrix-social-go-backend/db/v1"
)

func GetBookmark() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Locals("userID").(string)

		bookmarks, err := db.GetBookmark(userID)
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

		err := db.CreateBookmark(userID, postID)
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

		err := db.DeleteBookmark(userID, postID)
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
