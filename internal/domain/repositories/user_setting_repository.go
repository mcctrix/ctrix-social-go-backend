package repositories

import "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"

type UserSettingRepository interface {
	FindByID(id string) (*models.User_Settings, error)
	Update(settings *models.User_Settings) error
}
