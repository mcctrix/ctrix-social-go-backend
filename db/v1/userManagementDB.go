package v1

import (
	"encoding/json"
	"fmt"

	"github.com/mcctrix/ctrix-social-go-backend/models"
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
		// userProfile := &user_profile_data{}
		// user_auth_data, err := GetDataFromUserAuth(id)
		// if err != nil {
		// 	return nil, err
		// }
		// if len(fieldNames) > 0 {
		// 	if err = db.Table(tableName).Select(fieldNames).Where("id = ?", id).First(userProfile).Error; err != nil {
		// 		return nil, err
		// 	}
		// } else {
		// 	if err = db.Table(tableName).Where("id = ?", id).First(userProfile).Error; err != nil {
		// 		return nil, err
		// 	}
		// }
		// userProfile.Email = user_auth_data.Email
		// userProfile.Username = user_auth_data.Username

		// userData = userProfile

		// return userData, nil
	case "user_additional_info":
		userData = &models.User_Additional_Info{}
	case "user_settings":
		userData = &models.User_Settings{}
	case "user_data":
		userData = &models.User_Data{}
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
	case "user_data":
		userProfile = &models.User_Data{}
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
		Id string
	}

	data := base{Id: userid}

	if err = db.Table("user_profile").Create(data).Error; err != nil {
		return err
	}
	if err = db.Table("user_additional_info").Create(data).Error; err != nil {
		return err
	}
	if err = db.Table("user_data").Create(data).Error; err != nil {
		return err
	}
	userSetting := struct {
		Id          string
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
