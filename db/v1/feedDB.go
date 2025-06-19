package v1

import (
	"errors"

	"github.com/mcctrix/ctrix-social-go-backend/models"
)

type PostWithUserDetails struct {
	models.User_Post

	Username string `json:"username,omitempty"`

	Avatar         string `json:"avatar,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	VerifiedUser   bool   `json:"verified_user"`
	Bio            string `json:"bio,omitempty"`
}

func GetPostFeed(userID string, limit int) ([]PostWithUserDetails, error) {
	dbInstance, err := DBConnection()
	if err != nil {
		return nil, err
	}

	var userSettings models.User_Settings
	if err = dbInstance.Table("user_settings").Where("id = ?", userID).First(&userSettings).Error; err != nil {
		return nil, err
	}

	var postsWithDetails []PostWithUserDetails

	queryPost := dbInstance.Table("user_posts").
		Select("user_posts.*, user_auth.username, user_profile.avatar, user_profile.profile_picture, user_profile.verified_user, user_additional_info.bio").
		Joins("JOIN user_auth ON user_auth.id = user_posts.creator_id").
		Joins("JOIN user_profile ON user_profile.id = user_posts.creator_id").
		Joins("JOIN user_additional_info ON user_additional_info.id = user_posts.creator_id").
		Not("user_posts.id = ?", userSettings.Hide_post).
		Not("user_posts.creator_id = ?", userSettings.Block_user).
		Order("user_posts.created_at desc").
		Limit(limit).
		Find(&postsWithDetails)
	if err = queryPost.Error; err != nil {
		return nil, err
	}

	if queryPost.RowsAffected == 0 {
		return nil, errors.New("no posts found")
	}

	return postsWithDetails, nil
}

type followRecommendation struct {
	models.User_Profile

	Username string `json:"username,omitempty"`
}

func GetFollowRecommendation(currentUserID string, limit int) ([]followRecommendation, error) {
	dbInstance, err := DBConnection()
	if err != nil {
		return nil, err
	}

	var currentFollows []models.Follows
	if err = dbInstance.Table("follows").Where("follower_id = ?", currentUserID).Find(&currentFollows).Error; err != nil {
		return nil, err
	}
	var followingUsers []string
	for el := range currentFollows {
		followingUsers = append(followingUsers, currentFollows[el].Following_id)
	}
	followingUsers = append(followingUsers, currentUserID)

	var recommendation []followRecommendation
	if err = dbInstance.Table("user_profile").Select("user_auth.username, user_profile.*").Joins("JOIN user_auth ON user_auth.id = user_profile.id").Not("user_auth.id IN (?)", followingUsers).Limit(limit).Find(&recommendation).Error; err != nil {
		return nil, err
	}

	return recommendation, nil
}
