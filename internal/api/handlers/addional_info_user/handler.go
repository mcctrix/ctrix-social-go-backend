package addionalInfoUser

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
)

type AdditionalInfoUserHandler struct {
	additionalInfoUserService *services.AdditionalUserInfoService
}

func NewAdditionalInfoUserHandler(additionalInfoUserService *services.AdditionalUserInfoService) *AdditionalInfoUserHandler {
	return &AdditionalInfoUserHandler{
		additionalInfoUserService: additionalInfoUserService,
	}
}

func (h *AdditionalInfoUserHandler) GetAdditionalInfoUser(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	additionalInfo, err := h.additionalInfoUserService.GetAdditionalInfoByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("unable to fetch additional info user!")
	}
	return c.JSON(additionalInfo)
}

func (h *AdditionalInfoUserHandler) UpdateAdditionalInfoUser(c fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	rawData := c.BodyRaw()
	additionalInfo := &models.User_Additional_Info{}
	err := json.Unmarshal(rawData, additionalInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("unable to unmarshal additional info user!")
	}

	additionalInfo.Id = userID

	err = h.additionalInfoUserService.UpdateAdditionalInfo(additionalInfo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("unable to update additional info user!")
	}

	return c.SendString("Additional Info User Updated Successfully!")
}
