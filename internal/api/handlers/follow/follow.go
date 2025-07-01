package follow

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

type followHandler struct {
	followService *services.FollowService
}

func NewFollowHandler(followService *services.FollowService) *followHandler {
	return &followHandler{
		followService: followService,
	}
}

func (h *followHandler) FollowUser(c fiber.Ctx) error {
	followerID := c.Locals("userID").(string)
	followingID := c.Params("followingid")

	if err := h.followService.FollowUser(followerID, followingID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Followed Successfully!",
	})
}

func (h *followHandler) UnFollowUser(c fiber.Ctx) error {
	followerID := c.Locals("userID").(string)
	followingID := c.Params("followingid")

	if err := h.followService.UnFollowUser(followerID, followingID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Unfollowed Successfully!",
	})
}

func (h *followHandler) CheckFollowing(c fiber.Ctx) error {
	followerID := c.Locals("userID").(string)
	followingID := c.Params("followingid")

	isFollowing := h.followService.CheckFollowing(followerID, followingID)

	if isFollowing {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "following the user",
		})
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "not following the user",
	})
}

func (h *followHandler) CountFollowAndFollowing(c fiber.Ctx) error {
	userID := c.Params("userid")

	followAndFollowerCount, err := h.followService.GetFollowAndFollowing(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "data fetched successfully",
		"data":    followAndFollowerCount,
	})
}
