package repositories

import "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"

type FeedRepository interface {
	GetPostFeed(userID string, limit int) ([]models.PostWithUserDetails, error)
	GetFollowRecommendation(userID string, limit int) ([]models.FollowRecommendation, error)
}
