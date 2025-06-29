package repositories

import (
	"errors"
	"time"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database"
	"gorm.io/gorm"
)

func GetBookmark(userID string, limit int) ([]models.Bookmark, error) {
	dbInstance := database.GetDB()

	var bookmarks []models.Bookmark
	err := dbInstance.Table("bookmark").Where("user_id = ?", userID).Order("created_at desc").Limit(limit).Find(&bookmarks).Error
	if err != nil {
		return nil, err
	}
	return bookmarks, nil
}
func CreateBookmark(userID, postID string) error {
	dbInstance := database.GetDB()

	bookmark := models.Bookmark{
		User_id:    userID,
		Post_id:    postID,
		Created_at: time.Now(),
	}

	err := dbInstance.Create(&bookmark).Error
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

func DeleteBookmark(userID, postID string) error {
	dbInstance := database.GetDB()

	err := dbInstance.Table("bookmark").Where("user_id = ? AND post_id = ?", userID, postID).Delete(&models.Bookmark{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("bookmark does not exist")
		}
		return err
	}
	return nil
}
