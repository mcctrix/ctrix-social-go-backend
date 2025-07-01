package services

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

type AdditionalUserInfoService struct {
	userRepo repositories.AdditionalInfoRepository
}

func NewAdditionalService(userRepo repositories.AdditionalInfoRepository) *AdditionalUserInfoService {
	return &AdditionalUserInfoService{userRepo: userRepo}
}

func (s *AdditionalUserInfoService) GetAdditionalInfoByID(id string) (*models.User_Additional_Info, error) {
	additionalInfo, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return additionalInfo, nil
}
func (s *AdditionalUserInfoService) UpdateAdditionalInfo(additionalInfo *models.User_Additional_Info) error {
	return s.userRepo.Update(additionalInfo)
}
