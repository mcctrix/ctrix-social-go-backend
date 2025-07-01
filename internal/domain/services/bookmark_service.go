package services

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

type BookmarkService struct {
	bookmarkRepo repositories.BookmarkRepository
}

func NewBookmarkService(bookmarkRepo repositories.BookmarkRepository) *BookmarkService {
	return &BookmarkService{bookmarkRepo: bookmarkRepo}
}

func (s *BookmarkService) CreateBookmark(userID, postID string) error {
	return s.bookmarkRepo.CreateBookmark(userID, postID)
}

func (s *BookmarkService) DeleteBookmark(userID, postID string) error {
	return s.bookmarkRepo.DeleteBookmark(userID, postID)
}

func (s *BookmarkService) GetBookmark(userID string, limit int) ([]models.Bookmark, error) {
	data, err := s.bookmarkRepo.GetBookmark(userID, limit)
	if err != nil {
		return nil, err
	}

	return data, nil

}
