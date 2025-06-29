package repositories

import "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"

type ProfileRepository interface {
	FindByID(id string) (*models.User_Profile, error)
	Update(profile *models.User_Profile) error
}
