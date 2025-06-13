package v1

import "github.com/mcctrix/ctrix-social-go-backend/models"

func GetPostFeed(userID string) ([]models.User_Post, error) {
	dbInstance, err := DBConnection()
	if err != nil {
		return nil, err
	}

	var userSettings models.User_Settings
	if err = dbInstance.Table("user_settings").Where("id = ?", userID).First(&userSettings).Error; err != nil {
		return nil, err
	}

	var allPosts []models.User_Post

	if err = dbInstance.Table("user_posts").Not("id = ?", userSettings.Hide_post).
		Not("creator_id = ?", userSettings.Block_user).Order("created_at desc").Find(&allPosts).Error; err != nil {
		return nil, err
	}

	return allPosts, nil
}
