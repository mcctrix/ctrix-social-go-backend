package repositories

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"gorm.io/gorm"
)

type PostgresFollowRepository struct {
	db *gorm.DB
}

func NewPostgresFollowRepository(db *gorm.DB) *PostgresFollowRepository {
	return &PostgresFollowRepository{db: db}
}

func (r *PostgresFollowRepository) CreateFollow(follower_id, following_id string) error {
	follow := models.NewFollow(follower_id, following_id)
	if err := r.db.Model(&models.Follow{}).Create(follow).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresFollowRepository) IsFollowing(follower_id, following_id string) bool {
	var follow models.Follow
	res := r.db.Model(&models.Follow{}).Where("follower_id = ? AND following_id = ?", follower_id, following_id).First(&follow)
	return res.Error == nil
}

func (r *PostgresFollowRepository) UnFollow(follower_id, following_id string) error {
	if err := r.db.Table("follow").Where("follower_id = ? AND following_id = ?", follower_id, following_id).Delete(&models.Follow{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresFollowRepository) CountFollowAndFollowing(userID string) (*struct {
	followCount    int
	followingCount int
}, error) {
	var follows []models.Follow
	res := r.db.Model(&models.Follow{}).Select("follower_id, following_id").Where("follower_id = ? OR following_id = ?", userID, userID).Find(&follows)
	if err := res.Error; err != nil {
		return nil, err
	}
	followerCount := 0
	followingCount := 0
	for _, follow := range follows {
		if follow.Follower_id == userID {
			followerCount++
		} else {
			followingCount++
		}
	}
	result := &struct {
		followCount    int
		followingCount int
	}{
		followCount:    followerCount,
		followingCount: followingCount,
	}

	return result, nil
}
