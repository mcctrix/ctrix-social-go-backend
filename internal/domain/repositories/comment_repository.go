package repositories

import "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"

type CommentRepository interface {
	CreateComment(*models.User_post_Comments) error
	GetCommentByID(commentID string) (*models.User_post_Comments, error)
	GetCommentsByPostID(postID string, limit int) ([]models.User_post_Comments, error)
	GetCommentByUserID(userID string, limit int) ([]models.User_post_Comments, error)
	UpdateCommentByID(commentID string, updatedCommentData *models.User_post_Comments, userID string) error
	DeleteCommentByID(commentID string, userID string) error
}

type CommentReactionRepository interface {
	CreateCommentReaction(commentID string, userID string) error
	DeleteCommentReaction(commentID string, userID string) error
	GetCommentReactioCount(commentID string) (int, error)
}
