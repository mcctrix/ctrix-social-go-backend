package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mcctrix/ctrix-social-go-backend/models"
	"gorm.io/gorm"
)

func GetDataFromUserAuth(id string) (*user_auth_data, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var user_auth *user_auth_data = &user_auth_data{}
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

func GetUserProfileByID(id string) (*user_profile_data, error) {

	db, err := DBConnection()
	if err != nil {
		return nil, err
	}

	user_auth_data, err := GetDataFromUserAuth(id)
	if err != nil {
		return nil, err
	}

	var userProfile *user_profile_data = &user_profile_data{}
	if err = db.Table("user_profile").Where("id = ?", id).First(userProfile).Error; err != nil {
		return nil, err
	}
	userProfile.Email = user_auth_data.Email
	userProfile.Username = user_auth_data.Username
	return userProfile, nil
}

func SetUserProfileWithByteData(newProfileByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	var userProfile *models.User_Profile = &models.User_Profile{}

	if err = db.Table("user_profile").Where("id = ?", userID).First(userProfile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newProfile := &models.User_Profile{}
			err = json.Unmarshal(newProfileByte, newProfile)
			if err != nil {
				return err
			}
			newProfile.Id = userID
			err = createUserProfile(newProfile)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	json.Unmarshal(newProfileByte, userProfile)
	if err != nil {
		return err
	}
	db.Table("user_profile").Save(userProfile)

	return nil
}

type user_auth_data struct {
	Username string
	Email    string
}

func GetAdditionalInfoProfileByID(id string) (*models.User_Additional_Info, error) {

	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var userProfile *models.User_Additional_Info = &models.User_Additional_Info{}
	if err = db.Table("user_additional_info").Where("id = ?", id).First(userProfile).Error; err != nil {
		return nil, err
	}
	return userProfile, nil
}

func UpdateAdditionalUserProfileWithByteData(newProfileByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	var userProfile *models.User_Additional_Info = &models.User_Additional_Info{}
	if err = db.Table("user_additional_info").Where("id = ?", userID).First(userProfile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Additional Profile not found!")
		} else {
			fmt.Println(err)
		}
		return err
	}
	json.Unmarshal(newProfileByte, userProfile)
	if err != nil {
		fmt.Println("Error with the decode of json", err)
		return err
	}
	return db.Table("user_additional_info").Save(userProfile).Error
}

func createUserProfile(profileData *models.User_Profile) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}

	return db.Table("user_profile").Create(profileData).Error
}

func GetUserSettingsByID(id string) (*models.User_Settings, error) {

	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var userSettingsData *models.User_Settings = &models.User_Settings{}
	if err = db.Table("user_settings").Select("hide_post", "hide_story", "block_user", "show_online").Where("id = ?", id).First(userSettingsData).Error; err != nil {
		return nil, err
	}
	return userSettingsData, nil
}

func UpdateUserSettingsWithByteData(newProfileByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	var userSettings *models.User_Settings = &models.User_Settings{}

	if err = db.Table("user_settings").Where("id = ?", userID).First(userSettings).Error; err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(newProfileByte, userSettings)
	if err != nil {
		return err
	}
	return db.Table("user_settings").Save(userSettings).Error
}

func GetUserDataByID(id string) (*models.User_Data, error) {

	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var userData *models.User_Data = &models.User_Data{}
	if err = db.Table("user_data").Select("posts", "stories", "notes").Where("id = ?", id).First(userData).Error; err != nil {
		return nil, err
	}
	return userData, nil
}

func UpdateUserDataWithByteData(newUserDataByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	var userData *models.User_Data = &models.User_Data{}

	if err = db.Table("user_data").Where("id = ?", userID).First(userData).Error; err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(newUserDataByte, userData)
	if err != nil {
		return err
	}
	return db.Table("user_data").Save(userData).Error
}
