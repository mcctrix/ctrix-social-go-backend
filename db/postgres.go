package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mcctrix/ctrix-social-go-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func DBConnection() (*gorm.DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

	var host, username, password, dbname string = "", "", "", ""

	currentEnv := os.Getenv("APP_ENV")
	if currentEnv == "dev" {
		host = os.Getenv("postgresHostDev")
		dbname = os.Getenv("postgresDBDev")
		username = os.Getenv("postgresUsernameDev")
		password = os.Getenv("postgresPasswordDev")
	}

	if currentEnv == "production" {
		host = os.Getenv("postgresHostProd")
		dbname = os.Getenv("postgresDBProd")
		username = os.Getenv("postgresUsernameProd")
		password = os.Getenv("postgresPasswordProd")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, username, password, dbname)
	//Open the connection

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		return nil, err
	}
	dbInstance = db
	return dbInstance, nil
}

func CreateInitialDBStructure() {

	db, err := DBConnection()
	if err != nil {
		log.Fatal(err)
	}

	sqlFile, err := os.ReadFile("./sql/createTables.sql")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Exec(string(sqlFile)).Error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Initial Tables are Created Successfully!!!")
	}

}
func ResetDB() {
	db, err := DBConnection()
	if err != nil {
		fmt.Println("error here!1")
		log.Fatal("Error connecting to db: ", err)
	}
	// if err := db.Exec("DROP DATABASE Ctrix_Social_DB"); err != nil {
	// 	fmt.Println("error here!2")

	// 	log.Fatal(err)
	// }
	// if err := db.Exec("CREATE DATABASE IF NOT EXISTS Ctrix_Social_DB"); err != nil {
	// 	fmt.Println("error here!3")

	// 	log.Fatal(err)
	// }

	sqlFile, err := os.ReadFile("./sql/resetDB.sql")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Exec(string(sqlFile)).Error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB Resetted Successfully!!!")
	}

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

type user_auth_data struct {
	Username string
	Email    string
}

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
func CreateAdditionalUserProfileWithByteData(newProfileByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	var userProfile *models.User_Additional_Info = &models.User_Additional_Info{}
	err = json.Unmarshal(newProfileByte, userProfile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	userProfile.Id = userID

	return db.Table("user_additional_info").Create(userProfile).Error
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
func CreateUserSettingsWithByteData(userSettingsByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	var userSettings *models.User_Settings = &models.User_Settings{}
	if err = json.Unmarshal(userSettingsByte, userSettings); err != nil {
		fmt.Println(err)
		return err
	}
	userSettings.Id = userID

	return db.Table("user_settings").Create(userSettings).Error
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
func CreateUserDataWithByteData(newUserDataByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	var userData *models.User_Data = &models.User_Data{}
	if err = json.Unmarshal(newUserDataByte, userData); err != nil {
		fmt.Println(err)
		return err
	}

	return db.Table("user_data").Create(userData).Error
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
	return db.Table("user_profile").Save(userData).Error
}
