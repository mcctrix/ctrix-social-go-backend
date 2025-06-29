package services

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

type UserSettingService struct {
	userRepo repositories.UserSettingRepository
}

func (s *UserSettingService) GetUserSettingByID(id string) (*models.User_Settings, error) {
	settings, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return settings, nil
}
func (s *UserSettingService) UpdateUserSetting(settings *models.User_Settings) error {
	return s.userRepo.Update(settings)
}
