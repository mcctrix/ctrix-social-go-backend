package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"gorm.io/gorm"
)

type PostgresCommentRepository struct {
	db *gorm.DB
}

func NewPostgresCommentRepository(db *gorm.DB) *PostgresCommentRepository {
	return &PostgresCommentRepository{db: db}
}

func (r *PostgresCommentRepository) CreateComment(comment *models.User_post_Comments) error {
	if err := r.db.Model(&models.User_post_Comments{}).Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresCommentRepository) GetCommentByID(commentID string) (*models.User_post_Comments, error) {
	var comment *models.User_post_Comments = &models.User_post_Comments{}
	if err := r.db.Model(&models.User_post_Comments{}).Where("id = ?", commentID).First(comment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}
	return comment, nil
}

func (r *PostgresCommentRepository) GetCommentsByPostID(postID string, limit int) ([]models.User_post_Comments, error) {
	var comments []models.User_post_Comments
	if err := r.db.Model(&models.User_post_Comments{}).Where("post_id = ?", postID).Order("created_at desc").Limit(limit).Find(&comments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("comments not found")
		}
		return nil, err
	}
	return comments, nil
}

func (r *PostgresCommentRepository) GetCommentByUserID(userID string, limit int) ([]models.User_post_Comments, error) {
	var comments []models.User_post_Comments
	if err := r.db.Model(&models.User_post_Comments{}).Where("creator_id = ?", userID).Order("created_at desc").Limit(limit).Find(&comments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("comments not found")
		}
		return nil, err
	}
	return comments, nil
}

func (r *PostgresCommentRepository) UpdateCommentByID(commentID string, updatedCommentData *models.User_post_Comments, userID string) error {
	var comment *models.User_post_Comments = &models.User_post_Comments{}
	if err := r.db.Model(&models.User_post_Comments{}).Where("id = ?", commentID).First(comment).Error; err != nil {
		return err
	}
	if comment.Creator_id != userID {
		return errors.New("unauthorized to update comment")
	}
	comment.Updated_at = time.Now()
	comment.Content = updatedCommentData.Content
	comment.Giff = updatedCommentData.Giff
	if err := r.db.Save(comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresCommentRepository) DeleteCommentByID(commentID string, userID string) error {
	var comment *models.User_post_Comments = &models.User_post_Comments{}
	if err := r.db.Model(&models.User_post_Comments{}).Where("id = ?", commentID).First(comment).Error; err != nil {
		fmt.Println(err)
		return err
	}
	if comment.Creator_id != userID {
		return errors.New("unauthorized to delete comment")
	}
	if err := r.db.Model(&models.User_post_Comments{}).Delete(commentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("comment not found")
		}
		return err
	}
	return nil
}
