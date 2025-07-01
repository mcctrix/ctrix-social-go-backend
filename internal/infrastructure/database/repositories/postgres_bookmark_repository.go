package repositories

import (
	"errors"
	"time"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"gorm.io/gorm"
)

type PostgresBookmarkRepository struct {
	db *gorm.DB
}

func NewPostgresBookmarkRepository(db *gorm.DB) *PostgresBookmarkRepository {
	return &PostgresBookmarkRepository{db: db}
}

func (r *PostgresBookmarkRepository) GetBookmark(userID string, limit int) ([]models.Bookmark, error) {

	var bookmarks []models.Bookmark
	err := r.db.Model(&models.Bookmark{}).Where("user_id = ?", userID).Order("created_at desc").Limit(limit).Find(&bookmarks).Error
	if err != nil {
		return nil, err
	}
	return bookmarks, nil
}
func (r *PostgresBookmarkRepository) CreateBookmark(userID, postID string) error {

	bookmark := models.Bookmark{
		User_id:    userID,
		Post_id:    postID,
		Created_at: time.Now(),
	}

	err := r.db.Create(&bookmark).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("bookmark already exists")
		}
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return errors.New("user_id or post_id does not exist")
		}
		return err
	}
	return nil
}

func (r *PostgresBookmarkRepository) DeleteBookmark(userID, postID string) error {

	err := r.db.Model(&models.Bookmark{}).Where("user_id = ? AND post_id = ?", userID, postID).Delete(&models.Bookmark{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("bookmark does not exist")
		}
		return err
	}
	return nil
}
