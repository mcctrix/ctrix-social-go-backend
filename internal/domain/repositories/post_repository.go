package repositories

import "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"

type PostRepository interface {
	CreatePost(post *models.User_Post) error
	GetUserPostsByID(userID string, limit int) ([]models.User_Post, error)
	GetPostByID(postID string, userID string) (*models.User_Post, error)
	UpdatePost(postID string, updatedPost *models.User_Post, userID string) error
	DeletePost(postID string, userID string) error
	GetPostReactions(postID string) ([]models.User_Post_Like_Table, error)
	LikePost(postID string, userID string) error
	DislikePost(postID string, userID string) error
}
