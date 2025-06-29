package repositories

import "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"

type AdditionalInfoRepository interface {
	FindByID(id string) (*models.User_Additional_Info, error)
	Update(additionalInfo *models.User_Additional_Info) error
}
