package repositories

import (
	"errors"
	"fmt"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"gorm.io/gorm"
)

type PostgresFeedRepository struct {
	db *gorm.DB
}

func NewPostgresFeedRepository(db *gorm.DB) *PostgresFeedRepository {
	return &PostgresFeedRepository{db: db}
}

func (r *PostgresFeedRepository) GetPostFeed(userID string, limit int) ([]models.PostWithUserDetails, error) {
	var userSettings models.User_Settings
	if err := r.db.Table("user_settings").Where("id = ?", userID).First(&userSettings).Error; err != nil {
		return nil, err
	}

	var postsWithDetails []models.PostWithUserDetails
	queryPost := r.db.Table("user_posts").
		Select("user_posts.*, user_auth.username, user_profile.avatar, user_profile.profile_picture, user_profile.verified_user, user_additional_info.bio").
		Joins("JOIN user_auth ON user_auth.id = user_posts.creator_id").
		Joins("JOIN user_profile ON user_profile.id = user_posts.creator_id").
		Joins("JOIN user_additional_info ON user_additional_info.id = user_posts.creator_id").
		Not("user_posts.id = ?", userSettings.Hide_post).
		Not("user_posts.creator_id = ?", userSettings.Block_user).
		Order("user_posts.created_at desc").
		Limit(limit).
		Find(&postsWithDetails)
	if err := queryPost.Error; err != nil {
		fmt.Println("error while fetching post feed: ", err)
		return nil, err
	}
	if queryPost.RowsAffected == 0 {
		return nil, errors.New("no posts found")
	}

	for index := range postsWithDetails {
		var likes []models.User_Post_Like_Table
		query := r.db.Table("user_post_like")
		query.Select("user_id")
		query.Where("post_id = ?", postsWithDetails[index].Id)
		query.Find(&likes)
		if err := query.Error; err != nil {
			fmt.Println(err)
			return nil, err
		}
		postsWithDetails[index].LikesCount = len(likes)
	}

	for index := range postsWithDetails {
		liked, err := checkUserLikedPostGorm(r.db, postsWithDetails[index].Id, userID)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		postsWithDetails[index].IsLiked = liked
	}

	return postsWithDetails, nil
}

func (r *PostgresFeedRepository) GetFollowRecommendation(currentUserID string, limit int) ([]models.FollowRecommendation, error) {
	var currentFollows []models.Follow
	if err := r.db.Table("follow").Where("follower_id = ?", currentUserID).Find(&currentFollows).Error; err != nil {
		return nil, err
	}
	var followingUsers []string
	for el := range currentFollows {
		followingUsers = append(followingUsers, currentFollows[el].Following_id)
	}
	followingUsers = append(followingUsers, currentUserID)

	var recommendation []models.FollowRecommendation
	if err := r.db.Table("user_profile").Select("user_auth.username, user_profile.*").Joins("JOIN user_auth ON user_auth.id = user_profile.id").Not("user_auth.id IN (?)", followingUsers).Limit(limit).Find(&recommendation).Error; err != nil {
		return nil, err
	}

	return recommendation, nil
}

func checkUserLikedPostGorm(db *gorm.DB, postID string, userID string) (bool, error) {
	var likes []models.User_Post_Like_Table
	query := db.Table("user_post_like").Where("post_id = ? AND user_id = ?", postID, userID).Find(&likes)
	if err := query.Error; err != nil {
		return false, err
	}
	if query.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
