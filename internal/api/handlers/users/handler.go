package users

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

type ProfileHandler struct {
	profileService *services.ProfileService
}

func NewProfileHandler(profileService *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

func (h *ProfileHandler) GetUserProfileWithParam(c fiber.Ctx) error {
	userID := c.Params("userid")

	profile, err := h.profileService.GetProfileByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("unable to fetch user profile!")
	}
	return c.JSON(profile)
}

func (h *ProfileHandler) GetCurrentUserProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	profile, err := h.profileService.GetProfileByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("unable to fetch user profile!")
	}
	return c.JSON(profile)
}

func (h *ProfileHandler) UpdateCurrentUserProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	rawData := c.BodyRaw()
	profile := &models.User_Profile{}
	err := json.Unmarshal(rawData, profile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("unable to unmarshal user profile!")
	}

	profile.Id = userID

	err = h.profileService.UpdateProfile(profile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("unable to update user profile!")
	}

	return c.SendString("User Profile Updated Successfully!")
}
