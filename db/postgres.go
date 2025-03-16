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

func DBConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=ctrix password=6205 dbname=Ctrix_Social_DB sslmode=disable"
	//Open the connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateInitialDBStructure() {
	dsn := "host=localhost user=ctrix password=6205 dbname=Ctrix_Social_DB sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
func GetUserProfileByID(id string) (*models.User_Profile, error) {

	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	var userProfile *models.User_Profile = &models.User_Profile{}
	if err = db.Table("user_profile").Where("id = ?", id).First(userProfile).Error; err != nil {
		return nil, err
	}
	return userProfile, nil
}
func SetUserProfileWithByteData(newProfileByte []byte, userID string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	var userProfile *models.User_Profile = &models.User_Profile{}
	fmt.Println("old: ", userProfile)
	if err = db.Table("user_profile").Where("id = ?", userID).First(userProfile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newProfile := &models.User_Profile{}
			err = json.Unmarshal(newProfileByte, newProfile)
			if err != nil {
				fmt.Println("Error with the decode of json", err)
				return err
			}
			err = CreateUserProfile(newProfile)
			if err != nil {
				fmt.Println("Error Creating new profile:", err)
				return err
			}
		} else {
			return err
		}
	}
	json.Unmarshal(newProfileByte, userProfile)
	if err != nil {
		fmt.Println("Error with the decode of json", err)
		return err
	}
	db.Table("user_profile").Save(userProfile)

	return nil
}
func CreateUserProfile(profileData *models.User_Profile) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	fmt.Println(profileData.Id)
	db.Table("user_profile").Create(profileData)
	fmt.Println(profileData.Id)

	return nil
}
