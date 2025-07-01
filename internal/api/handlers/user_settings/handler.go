package userSetting

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

type userSettingsHandler struct {
	userSettingsService *services.UserSettingService
}

func NewUserSettingsHandler(userSettingsService *services.UserSettingService) *userSettingsHandler {
	return &userSettingsHandler{
		userSettingsService: userSettingsService,
	}
}

func (h *userSettingsHandler) GetUserSettings(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	settings, err := h.userSettingsService.GetUserSettingByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("unable to fetch user settings!")
	}
	return c.JSON(settings)
}

func (h *userSettingsHandler) UpdateUserSettings(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	rawData := c.BodyRaw()
	settings := &models.User_Settings{}
	err := json.Unmarshal(rawData, settings)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("unable to unmarshal user settings!")
	}

	settings.Id = userID

	err = h.userSettingsService.UpdateUserSetting(settings)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("unable to update user settings!")
	}

	return c.SendString("User Settings Updated Successfully!")
}
