package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database"
)

type dataInterface interface{}

func GetDataFromUserAuth(id string) (*models.User, error) {
	db := database.GetDB()

	var user_auth *models.User = &models.User{}
	if err := db.Table("user_auth").Select("username", "email").Where("id = ?", id).First(user_auth).Error; err != nil {
		return nil, err
	}

	return user_auth, nil
}

func GetUserData(id string, tableName string, fieldNames []string) (dataInterface, error) {

	db := database.GetDB()

	var userData dataInterface

	switch tableName {
	case "user_auth":
		userData = &models.User{}
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
		if err := db.Table(tableName).Select(fieldNames).Where("id = ?", id).First(userData).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Table(tableName).Where("id = ?", id).First(userData).Error; err != nil {
			return nil, err
		}
	}

	return userData, nil
}

func UpdateTableWithByteData(NewProfileData []byte, userID string, tableName string) error {
	db := database.GetDB()

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

	if err := db.Table(tableName).Where("id = ?", userID).First(&userProfile).Error; err != nil {
		return err
	}

	if err := json.Unmarshal(NewProfileData, userProfile); err != nil {
		return err
	}
	return db.Table(tableName).Save(userProfile).Error
}

func InitNewUser(userid string) error {
	db := database.GetDB()

	type base struct {
		Id string `gorm:"primaryKey"`
	}

	data := base{Id: userid}

	if err := db.Table("user_profile").Create(data).Error; err != nil {
		return err
	}
	if err := db.Table("user_additional_info").Create(data).Error; err != nil {
		return err
	}
	userSetting := struct {
		Id          string `gorm:"primaryKey"`
		Show_online bool
	}{
		Id:          userid,
		Show_online: true,
	}

	if err := db.Table("user_settings").Create(userSetting).Error; err != nil {
		return err
	}
	return nil
}
