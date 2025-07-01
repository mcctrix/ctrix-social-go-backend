package services

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/repositories"
)

type FollowService struct {
	followRepo repositories.FollowRepository
}

func NewFollowService(followRepo repositories.FollowRepository) *FollowService {
	return &FollowService{followRepo: followRepo}
}

func (s *FollowService) FollowUser(follower_id, following_id string) error {
	return s.followRepo.CreateFollow(follower_id, following_id)
}

func (s *FollowService) UnFollowUser(follower_id, following_id string) error {
	return s.followRepo.UnFollow(follower_id, following_id)
}

func (s *FollowService) CheckFollowing(follower_id, following_id string) (*models.Follow, error) {
	data, err := s.followRepo.IsFollowing(follower_id, following_id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type followAndFollowerCount struct {
	FollowerCount  int `json:"follower_count"`
	FollowingCount int `json:"following_count"`
}

func (s *FollowService) GetFollowAndFollowing(userID string) (*followAndFollowerCount, error) {
	followCount, followingCount, err := s.followRepo.CountFollowAndFollowing(userID)
	if err != nil {
		return nil, err
	}

	return &followAndFollowerCount{
		FollowerCount:  followCount,
		FollowingCount: followingCount,
	}, nil
}
