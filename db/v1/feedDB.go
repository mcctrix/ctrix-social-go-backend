package v1

import (
	"errors"

	"github.com/mcctrix/ctrix-social-go-backend/models"
)

type PostWithUserDetails struct {
	// Fields from User_Post (using anonymous embedding for convenience, or list them individually)
	models.User_Post

	// Fields from User_Auth
	Username string `json:"username,omitempty"`

	// Fields from User_Profile
	Avatar         string `json:"avatar,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	VerifiedUser   bool   `json:"verified_user"`

	// Fields from User_Additional_Info
	Bio string `json:"bio,omitempty"`
}

func GetPostFeed(userID string) ([]PostWithUserDetails, error) {
	dbInstance, err := DBConnection()
	if err != nil {
		return nil, err
	}

	var userSettings models.User_Settings
	if err = dbInstance.Table("user_settings").Where("id = ?", userID).First(&userSettings).Error; err != nil {
		return nil, err
	}

	/*
		veriedUser ?
		poster_username
		poster_avatar
		poster_profile_picture

	*/

	var postsWithDetails []PostWithUserDetails

	queryPost := dbInstance.Table("user_posts").
		Select("user_posts.*, user_auth.username, user_profile.avatar, user_profile.profile_picture, user_profile.verified_user, user_additional_info.bio").
		Joins("JOIN user_auth ON user_auth.id = user_posts.creator_id").
		Joins("JOIN user_profile ON user_profile.id = user_posts.creator_id").
		Joins("JOIN user_additional_info ON user_additional_info.id = user_posts.creator_id").
		Not("user_posts.id = ?", userSettings.Hide_post).
		Not("user_posts.creator_id = ?", userSettings.Block_user).
		Order("user_posts.created_at desc").
		Find(&postsWithDetails)
	if err = queryPost.Error; err != nil {
		return nil, err
	}

	if queryPost.RowsAffected == 0 {
		return nil, errors.New("no posts found")
	}

	return postsWithDetails, nil
}
