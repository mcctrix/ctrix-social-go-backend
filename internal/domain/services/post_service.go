package services

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

type PostService struct {
	postRepo repositories.PostRepository
}

func NewPostService(postRepo repositories.PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) CreatePost(post *models.User_Post) error {
	return s.postRepo.CreatePost(post)
}

func (s *PostService) GetUserPostsByID(userID string, limit int) ([]models.User_Post, error) {
	return s.postRepo.GetUserPostsByID(userID, limit)
}

func (s *PostService) GetPostByID(postID string, userID string) (*models.User_Post, error) {
	return s.postRepo.GetPostByID(postID, userID)
}

func (s *PostService) UpdatePost(postID string, updatedPost *models.User_Post, userID string) error {
	return s.postRepo.UpdatePost(postID, updatedPost, userID)
}

func (s *PostService) DeletePost(postID string, userID string) error {
	return s.postRepo.DeletePost(postID, userID)
}

func (s *PostService) GetPostReactions(postID string) ([]models.User_Post_Like_Table, error) {
	return s.postRepo.GetPostReactions(postID)
}

func (s *PostService) TogglePostLike(postID string, userID string, liked bool, likeType string) error {
	return s.postRepo.TogglePostLike(postID, userID, liked, likeType)
} 