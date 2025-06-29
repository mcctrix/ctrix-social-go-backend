package services

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

type ProfileService struct {
	userRepo repositories.ProfileRepository
}

func (s *ProfileService) GetProfileByID(id string) (*models.User_Profile, error) {
	profile, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
func (s *ProfileService) UpdateProfile(profile *models.User_Profile) error {
	return s.userRepo.Update(profile)
}
