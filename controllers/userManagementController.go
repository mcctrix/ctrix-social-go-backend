package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	db "github.com/mcctrix/ctrix-social-go-backend/db/v1"
	"github.com/mcctrix/ctrix-social-go-backend/models"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
)

func ProfileSetup() fiber.Handler {
	return func(c fiber.Ctx) error {
		// avatar,bio,gender,dob
		dataInterface := &struct {
			Avatar string    `json:"avatar,omitempty"`
			Dob    time.Time `json:"dob,omitempty"`
			Bio    string    `json:"bio,omitempty"`
			Gender string    `json:"gender,omitempty"`
		}{}

		rawForm, err := utils.ClearStruct(dataInterface, c.BodyRaw())
		if err != nil {
			fmt.Println("Error Marshalling the data: ", err)
			return fiber.ErrInternalServerError
		}

		err = db.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_additional_info")
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}

		err = db.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_profile")
		if err != nil {
			fmt.Println("Error Setting the user profile: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Profile Setup Page")
	}
}

func GetCurrentUserProfile() fiber.Handler {
	return func(c fiber.Ctx) error {
		profile, err := db.GetUserData(c.Locals("userID").(string), "user_profile", []string{"first_name", "last_name", "avatar", "last_seen", "post_count", "followers", "followings"})
		if err != nil {
			fmt.Println("unable to fetch profile: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch user profile!")
		}
		userAuth, err := db.GetUserData(c.Locals("userID").(string), "user_auth", []string{"email", "username"})
		if err != nil {
			fmt.Println("unable to fetch user Auth: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch user auth!")
		}
		merged, err := utils.MergeStructs(profile, userAuth)
		if err != nil {
			fmt.Println("unable to merge structs: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to merge structs!")
		}
		return c.JSON(merged)
	}
}

func GetUserProfileWithParam() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Params("userid")
		profile, err := db.GetUserData(userID, "user_profile", []string{"first_name", "last_name", "avatar", "last_seen", "post_count", "followers", "followings"})
		if err != nil {
			fmt.Println("unable to fetch profile: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch user profile!")
		}
		userAuth, err := db.GetUserData(userID, "user_auth", []string{"email", "username"})
		if err != nil {
			fmt.Println("unable to fetch user Auth: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch user auth!")
		}
		merged, err := utils.MergeStructs(profile, userAuth)
		if err != nil {
			fmt.Println("unable to merge structs: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to merge structs!")
		}
		return c.JSON(merged)
	}
}

func UpdateCurrentUserProfile() fiber.Handler {
	return func(c fiber.Ctx) error {
		dataInterface := &struct {
			First_name      string `json:"first_name,omitempty"`
			Last_name       string `json:"last_name,omitempty"`
			Profile_picture string `json:"profile_profile,omitempty"`
		}{}

		rawForm, err := utils.ClearStruct(dataInterface, c.BodyRaw())
		if err != nil {
			fmt.Println("Error Marshalling the data: ", err)
			return fiber.ErrInternalServerError
		}

		err = db.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_profile")
		if err != nil {
			fmt.Println("Error Setting the profile: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Profile is updated successfully!")
	}
}

func GetAdditionalUserInfo() fiber.Handler {
	return func(c fiber.Ctx) error {

		additional_profile, err := db.GetUserData(c.Locals("userID").(string), "user_additional_info", []string{"hobbies", "family_members", "relation_status", "dob", "bio", "gender"})
		if err != nil {
			fmt.Println("error while fetching additional profile: ", err)
			return c.Status(500).SendString("unable to fetch user additional profile!")
		}

		return c.JSON(additional_profile)
	}
}

func UpdateAdditionalUserInfo() fiber.Handler {
	return func(c fiber.Ctx) error {
		dataInterface := &struct {
			Hobbies         models.StringArray `json:"hobbies,omitempty" gorm:"type:text[]"`
			Family_members  models.StringArray `json:"family_members,omitempty" gorm:"type:text[]"`
			Relation_status string             `json:"relation_status,omitempty"`
			Avatar          string             `json:"avatar,omitempty"`
			Dob             time.Time          `json:"dob,omitempty"`
			Bio             string             `json:"bio,omitempty"`
			Gender          string             `json:"gender,omitempty"`
		}{}

		rawForm, err := utils.ClearStruct(dataInterface, c.BodyRaw())
		if err != nil {
			fmt.Println("Error Marshalling the data: ", err)
			return fiber.ErrInternalServerError
		}

		err = db.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_additional_info")
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}
		return c.SendString("User Additional profile updated Successfully!")
	}
}

func GetUserSettings() fiber.Handler {
	return func(c fiber.Ctx) error {
		additional_profile, err := db.GetUserData(c.Locals("userID").(string), "user_settings", []string{"hide_post", "hide_story", "block_user", "show_online"})
		if err != nil {
			fmt.Println("error while fetching additional profile: ", err)
			return c.Status(500).SendString("unable to fetch user settings!")
		}

		return c.JSON(additional_profile)
	}
}

func UpdateUserSettings() fiber.Handler {
	return func(c fiber.Ctx) error {
		dataInterface := &struct {
			Hide_post   models.StringArray `json:"hide_post,omitempty" gorm:"type:text[]"`
			Hide_story  models.StringArray `json:"hide_story,omitempty" gorm:"type:text[]"`
			Block_user  models.StringArray `json:"block_user,omitempty" gorm:"type:text[]"`
			Show_online bool               `json:"show_online" gorm:"type:text[]"`
		}{}

		rawForm, err := utils.ClearStruct(dataInterface, c.BodyRaw())
		if err != nil {
			fmt.Println("Error Marshalling the data: ", err)
			return fiber.ErrInternalServerError
		}

		err = db.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_settings")
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}
		return c.SendString("User Additional profile updated Successfully!")
	}
}

func GetUserData() fiber.Handler {
	return func(c fiber.Ctx) error {

		additional_profile, err := db.GetUserData(c.Locals("userID").(string), "user_data", []string{"posts", "stories", "notes"})
		if err != nil {
			fmt.Println("error while fetching additional profile: ", err)
			return c.Status(500).SendString("unable to fetch user data!")
		}

		return c.JSON(additional_profile)
	}
}

func UpdateUserData() fiber.Handler {
	return func(c fiber.Ctx) error {

		dataInterface := &struct {
			Posts   models.StringArray `json:"posts,omitempty" gorm:"type:text[]"`
			Stories models.StringArray `json:"stories,omitempty" gorm:"type:text[]"`
			Notes   models.StringArray `json:"notes,omitempty" gorm:"type:text[]"`
		}{}

		rawForm, err := utils.ClearStruct(dataInterface, c.BodyRaw())
		if err != nil {
			fmt.Println("Error Marshalling the data: ", err)
			return fiber.ErrInternalServerError
		}

		err = db.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_data")
		if err != nil {
			fmt.Println("Error Setting the user data: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("User Data updated Successfully!")
	}
}
