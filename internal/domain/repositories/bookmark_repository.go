package repositories

import "github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"

type BookmarkRepository interface {
	GetBookmark(userID string, limit int) ([]models.Bookmark, error)
	CreateBookmark(userID, postID string) error
	DeleteBookmark(userID, postID string) error
}
