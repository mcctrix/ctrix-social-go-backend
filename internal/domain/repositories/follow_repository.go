package repositories

import "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"

type FollowRepository interface {
	CreateFollow(follower_id, following_id string) error
	IsFollowing(follower_id, following_id string) (*models.Follow, error)
	UnFollow(follower_id, following_id string) error
	CountFollowAndFollowing(userID string) (int, int, error)
}
