package repositories

import (
	"errors"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"gorm.io/gorm"
)

type PostgresCommentReactionRepository struct {
	db *gorm.DB
}

func NewPostgresCommentReactionRepository(db *gorm.DB) *PostgresCommentReactionRepository {
	return &PostgresCommentReactionRepository{db: db}
}

func (r *PostgresCommentReactionRepository) CreateCommentReaction(commentID string, userID string) error {
	react := &models.User_post_comment_like{User_id: userID, Comment_id: commentID}
	if err := r.db.Model(&models.User_post_comment_like{}).Create(react).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("already liked the comment")
		}
		return err
	}
	return nil
}

func (r *PostgresCommentReactionRepository) DeleteCommentReaction(commentID string, userID string) error {
	if err := r.db.Model(&models.User_post_comment_like{}).Where("comment_id = ? AND user_id = ?", commentID, userID).Delete(&models.User_post_comment_like{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresCommentReactionRepository) GetCommentReactioCount(commentID string) (int, error) {
	var likes []models.User_post_comment_like

	query := r.db.Model(&models.User_post_comment_like{}).Where("comment_id = ?", commentID).Find(&likes)
	if err := query.Error; err != nil {
		return 0, err
	}

	if query.RowsAffected == 0 {
		return 0, nil
	}
	return len(likes), nil
}
