package repositories

import (
	"errors"
	"fmt"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"gorm.io/gorm"
)

type PostgresProfileRepository struct {
	db *gorm.DB
}

func NewPostgresProfileRepository(DB *gorm.DB) *PostgresProfileRepository {
	return &PostgresProfileRepository{db: DB}
}

func (r *PostgresProfileRepository) FindByID(id string) (*models.User_Profile, error) {
	var profile *models.User_Profile = new(models.User_Profile)
	query := r.db.Model(&models.User_Profile{}).Select("first_name, last_name, avatar, profile_picture, last_seen, verified_user").Where("id = ?", id).First(profile)
	if err := query.Error; err != nil {
		fmt.Println("error in getting profile: ", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to find profile: %w", err)
		}
		return nil, err
	}
	return profile, nil
}

func (r *PostgresProfileRepository) Update(profile *models.User_Profile) error {
	query := r.db.Model(&models.User_Profile{}).Where("id = ?", profile.Id).Updates(profile)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}
