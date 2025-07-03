package services

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

type FeedService struct {
	feedRepo repositories.FeedRepository
}

func NewFeedService(feedRepo repositories.FeedRepository) *FeedService {
	return &FeedService{feedRepo: feedRepo}
}

func (s *FeedService) GetPostFeed(userID string, limit int) ([]models.PostWithUserDetails, error) {
	return s.feedRepo.GetPostFeed(userID, limit)
}

func (s *FeedService) GetFollowRecommendation(userID string, limit int) ([]models.FollowRecommendation, error) {
	return s.feedRepo.GetFollowRecommendation(userID, limit)
}
