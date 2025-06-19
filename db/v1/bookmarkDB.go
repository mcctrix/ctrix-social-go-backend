package v1

import (
	"errors"
	"time"

	"github.com/mcctrix/ctrix-social-go-backend/models"
	"gorm.io/gorm"
)

func GetBookmark(userID string) ([]models.Bookmark, error) {
	dbInstance, err := DBConnection()
	if err != nil {
		return nil, err
	}

	var bookmarks []models.Bookmark
	err = dbInstance.Table("bookmark").Where("user_id = ?", userID).Find(&bookmarks).Error
	if err != nil {
		return nil, err
	}
	return bookmarks, nil
}
func CreateBookmark(userID, postID string) error {
	dbInstance, err := DBConnection()
	if err != nil {
		return err
	}

	bookmark := models.Bookmark{
		User_id:    userID,
		Post_id:    postID,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	err = dbInstance.Create(&bookmark).Error
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
	dbInstance, err := DBConnection()
	if err != nil {
		return err
	}

	err = dbInstance.Table("bookmark").Where("user_id = ? AND post_id = ?", userID, postID).Delete(&models.Bookmark{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("bookmark does not exist")
		}
		return err
	}
	return nil
}
