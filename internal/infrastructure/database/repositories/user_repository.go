package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"gorm.io/gorm"
)

type dataInterface interface{}

func GetDataFromUserAuth(id string) (*models.User_Auth, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var user_auth *models.User_Auth = &models.User_Auth{}
	if err = db.Table("user_auth").Select("username", "email").Where("id = ?", id).First(user_auth).Error; err != nil {
		return nil, err
	}

	return user_auth, nil
}

type user_profile_data struct {
	models.User_Profile
	Email    string `json:"email"`
	Username string `json:"username"`
}

func GetUserData(id string, tableName string, fieldNames []string) (dataInterface, error) {

	db, err := DBConnection()
	if err != nil {
		return nil, err
	}

	var userData dataInterface

	switch tableName {
	case "user_auth":
		userData = &models.User_Auth{}
	case "user_profile":
		userData = &models.User_Profile{}
	case "user_additional_info":
		userData = &models.User_Additional_Info{}
	case "user_settings":
		userData = &models.User_Settings{}
	default:
		return nil, fmt.Errorf("unsupported table name for get: %s", tableName)
	}

	if len(fieldNames) > 0 {
		if err = db.Table(tableName).Select(fieldNames).Where("id = ?", id).First(userData).Error; err != nil {
			return nil, err
		}
	} else {
		if err = db.Table(tableName).Where("id = ?", id).First(userData).Error; err != nil {
			return nil, err
		}
	}

	return userData, nil
}

func UpdateTableWithByteData(NewProfileData []byte, userID string, tableName string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	var userProfile dataInterface

	switch tableName {
	case "user_profile":
		userProfile = &models.User_Profile{}
	case "user_additional_info":
		userProfile = &models.User_Additional_Info{}
	case "user_settings":
		userProfile = &models.User_Settings{}
	default:
		return fmt.Errorf("unsupported table name for update: %s", tableName)
	}

	if err = db.Table(tableName).Where("id = ?", userID).First(&userProfile).Error; err != nil {
		return err
	}

	if err = json.Unmarshal(NewProfileData, userProfile); err != nil {
		return err
	}
	return db.Table(tableName).Save(userProfile).Error
}

func InitNewUser(userid string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	type base struct {
		Id string `gorm:"primaryKey"`
	}

	data := base{Id: userid}

	if err = db.Table("user_profile").Create(data).Error; err != nil {
		return err
	}
	if err = db.Table("user_additional_info").Create(data).Error; err != nil {
		return err
	}
	userSetting := struct {
		Id          string `gorm:"primaryKey"`
		Show_online bool
	}{
		Id:          userid,
		Show_online: true,
	}

	if err = db.Table("user_settings").Create(userSetting).Error; err != nil {
		return err
	}
	return nil
}

func FollowUser(follow_id string, following_id string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	if err = db.Table("follows").Create(&models.Follows{
		Follower_id:  follow_id,
		Following_id: following_id,
		Created_at:   time.Now(),
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return errors.New("following user not found")
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("user already followed")
		}

		return err
	}

	return nil
}

func UnfollowUser(follow_id string, following_id string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	if err = db.Table("follows").Where("follower_id = ? AND following_id = ?", follow_id, following_id).Delete(&models.Follows{}).Error; err != nil {
		return err
	}

	return nil
}

func CheckFollowing(follow_id string, following_id string) (*models.Follows, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}

	var follows *models.Follows
	res := db.Table("follows").Select("created_at").Where("follower_id = ? AND following_id = ?", follow_id, following_id).Find(&follows)
	if err = res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("Not Following")
	}

	return follows, nil
}

type FollowAndFollowerCount struct {
	FollowerCount  int `json:"follower_count"`
	FollowingCount int `json:"following_count"`
}

func GetFollowAndFollowing(userID string) (*FollowAndFollowerCount, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}

	var follows []models.Follows
	res := db.Table("follows").Select("follower_id, following_id").Where("follower_id = ? OR following_id = ?", userID, userID).Find(&follows)
	if err = res.Error; err != nil {
		return nil, err
	}
	data := &FollowAndFollowerCount{}
	for _, follow := range follows {
		if follow.Follower_id == userID {
			data.FollowerCount++
		} else {
			data.FollowingCount++
		}
	}

	return data, nil
}
