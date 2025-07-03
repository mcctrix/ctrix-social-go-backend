package services

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

type CommentReactionService struct {
	commentReactionRepository repositories.CommentReactionRepository
}

func NewCommentReactionService(commentReactionRepository repositories.CommentReactionRepository) *CommentReactionService {
	return &CommentReactionService{commentReactionRepository: commentReactionRepository}
}

func (s *CommentReactionService) CreateCommentReaction(commentID string, userID string) error {
	return s.commentReactionRepository.CreateCommentReaction(commentID, userID)
}

func (s *CommentReactionService) DeleteCommentReaction(commentID string, userID string) error {
	return s.commentReactionRepository.DeleteCommentReaction(commentID, userID)
}

func (s *CommentReactionService) GetCommentReactioCount(commentID string) (int, error) {
	return s.commentReactionRepository.GetCommentReactioCount(commentID)
}
