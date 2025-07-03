package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"gorm.io/gorm"
)

type PostgresPostRepository struct {
	db *gorm.DB
}

func NewPostgresPostRepository(db *gorm.DB) *PostgresPostRepository {
	return &PostgresPostRepository{db: db}
}

func (r *PostgresPostRepository) CreatePost(post *models.User_Post) error {
	return r.db.Model(&models.User_Post{}).Create(post).Error
}

func (r *PostgresPostRepository) GetUserPostsByID(userID string, limit int) ([]models.User_Post, error) {
	var posts []models.User_Post
	err := r.db.Model(&models.User_Post{}).Where("creator_id = ?", userID).Order("created_at desc").Limit(limit).Find(&posts).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no posts found")
		}
		return nil, err
	}
	return posts, nil
}

func (r *PostgresPostRepository) GetPostByID(postID string, userID string) (*models.User_Post, error) {
	var post models.User_Post
	err := r.db.Model(&models.User_Post{}).Where("id = ?", postID).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostgresPostRepository) UpdatePost(postID string, updatedPost *models.User_Post, userID string) error {
	var post models.User_Post
	if err := r.db.Model(&models.User_Post{}).Where("id = ? AND creator_id = ?", postID, userID).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return err
	}
	post.Text_content = updatedPost.Text_content
	post.Media_attached = updatedPost.Media_attached
	post.Updated_at = time.Now()
	return r.db.Save(&post).Error
}

func (r *PostgresPostRepository) DeletePost(postID string, userID string) error {
	result := r.db.Model(&models.User_Post{}).Where("id = ? AND creator_id = ?", postID, userID).Delete(&models.User_Post{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no post found or unauthorized to delete")
	}
	return nil
}

func (r *PostgresPostRepository) GetPostReactions(postID string) ([]models.User_Post_Like_Table, error) {
	var reacts []models.User_Post_Like_Table
	err := r.db.Model(&models.User_Post_Like_Table{}).Where("post_id = ?", postID).Find(&reacts).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no reactions found")
		}
		return nil, err
	}
	return reacts, nil
}

func (r *PostgresPostRepository) LikePost(postID string, userID string) error {
	post_like_data := models.User_Post_Like_Table{User_id: userID, Post_id: postID}
	if err := r.db.Table("user_post_like").Create(post_like_data).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *PostgresPostRepository) DislikePost(postID string, userID string) error {
	if err := r.db.Table("user_post_like").Where("user_id = ? AND post_id = ?", userID, postID).Delete(&models.User_Post_Like_Table{}).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
