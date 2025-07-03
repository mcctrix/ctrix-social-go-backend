package feed

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
)

type FeedHandler struct {
	feedService *services.FeedService
}

func NewFeedHandler(feedService *services.FeedService) *FeedHandler {
	return &FeedHandler{feedService: feedService}
}

func (h *FeedHandler) GetFeed(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	limit := utils.QueryLimit(c.Query("limit", "5"))
	posts, err := h.feedService.GetPostFeed(userID, limit)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "unable to fetch feed!"})
	}
	return c.JSON(posts)
}

func (h *FeedHandler) GetFollowRecommendation(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	limit := utils.QueryLimit(c.Query("limit", "5"))
	recommendation, err := h.feedService.GetFollowRecommendation(userID, limit)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "unable to fetch follow recommendation!"})
	}
	return c.JSON(recommendation)
}
