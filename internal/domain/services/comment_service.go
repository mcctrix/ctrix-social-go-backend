package services

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

type CommentService struct {
	commentRepo repositories.CommentRepository
}

func NewCommentService(commentRepo repositories.CommentRepository) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}

func (s *CommentService) CreateComment(comment *models.User_post_Comments) error {
	return s.commentRepo.CreateComment(comment)
}

func (s *CommentService) GetCommentByID(commentID string) (*models.User_post_Comments, error) {
	return s.commentRepo.GetCommentByID(commentID)
}

func (s *CommentService) GetCommentsByPostID(postID string, limit int) ([]models.User_post_Comments, error) {
	return s.commentRepo.GetCommentsByPostID(postID, limit)
}

func (s *CommentService) GetCommentByUserID(userID string, limit int) ([]models.User_post_Comments, error) {
	return s.commentRepo.GetCommentByUserID(userID, limit)
}

func (s *CommentService) UpdateCommentByID(commentID string, updatedCommentData *models.User_post_Comments, userID string) error {
	return s.commentRepo.UpdateCommentByID(commentID, updatedCommentData, userID)
}

func (s *CommentService) DeleteCommentByID(commentID string, userID string) error {
	return s.commentRepo.DeleteCommentByID(commentID, userID)
}
