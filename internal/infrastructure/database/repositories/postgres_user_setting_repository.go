package repositories

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"gorm.io/gorm"
)

type PostgresUserSettingRepository struct {
	db *gorm.DB
}

func NewPostgresUserSettingRepository(db *gorm.DB) *PostgresUserSettingRepository {
	return &PostgresUserSettingRepository{db: db}
}

func (r *PostgresUserSettingRepository) FindByID(id string) (*models.User_Settings, error) {
	var settings *models.User_Settings
	query := r.db.Model(&settings).Where("id = ?", id).First(&settings)
	if query.Error != nil {
		return nil, query.Error
	}
	return settings, nil
}

func (r *PostgresUserSettingRepository) Update(settings *models.User_Settings) error {
	query := r.db.Model(&settings).Where("id = ?", settings.Id).Updates(settings)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
