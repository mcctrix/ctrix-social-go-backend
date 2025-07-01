package bookmarks

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
)

type bookmarkHandler struct {
	bookmarkService *services.BookmarkService
}

func NewBookmarkHandler(bookmarkService *services.BookmarkService) *bookmarkHandler {
	return &bookmarkHandler{
		bookmarkService: bookmarkService,
	}
}

func (h *bookmarkHandler) GetBookmark(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	limit := utils.QueryLimit(c.Query("limit"))

	bookmarks, err := h.bookmarkService.GetBookmark(userID, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(bookmarks)
}

func (h *bookmarkHandler) CreateBookmark(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	postID := c.Params("postID")

	err := h.bookmarkService.CreateBookmark(userID, postID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "bookmark created successfully",
	})
}

func (h *bookmarkHandler) DeleteBookmark(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	postID := c.Params("postID")

	err := h.bookmarkService.DeleteBookmark(userID, postID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "bookmark deleted successfully",
	})
}
