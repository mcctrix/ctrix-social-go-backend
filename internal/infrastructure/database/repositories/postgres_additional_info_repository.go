package repositories

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"gorm.io/gorm"
)

type PostgresAdditionalInfoRepository struct {
	db *gorm.DB
}

func NewPostgresAdditionalInfoRepository(db *gorm.DB) *PostgresAdditionalInfoRepository {
	return &PostgresAdditionalInfoRepository{db: db}
}

func (r *PostgresAdditionalInfoRepository) FindByID(id string) (*models.User_Additional_Info, error) {
	var additionalInfo *models.User_Additional_Info = new(models.User_Additional_Info)
	query := r.db.Model(&models.User_Additional_Info{}).Select("hobbies, relation_status, dob, bio, gender").Where("id = ?", id).First(additionalInfo)
	if query.Error != nil {
		return nil, query.Error
	}
	return additionalInfo, nil
}

func (r *PostgresAdditionalInfoRepository) Update(additionalInfo *models.User_Additional_Info) error {
	query := r.db.Model(&models.User_Additional_Info{}).Where("id = ?", additionalInfo.Id).Updates(additionalInfo)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
