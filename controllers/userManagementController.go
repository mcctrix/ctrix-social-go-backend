package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/db"
	"github.com/mcctrix/ctrix-social-go-backend/models"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
)

func GetCurrentUserProfile() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}
		profile, err := db.GetUserProfileByID(userID)
		if err != nil {
			fmt.Println("unable to fetch profile: ", err)
			return c.Status(500).SendString("unable to fetch user profile!")
		}
		return c.JSON(profile)
	}
}
func GetUserProfileWithParam() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Params("userid")
		profile, err := db.GetUserProfileByID(userID)
		if err != nil {
			fmt.Println("unable to fetch profile: ", err)
			return c.Status(500).SendString("unable to fetch user profile!")
		}
		return c.JSON(profile)
	}
}
func SetCurrentUserProfile() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}
		newProfile := &models.User_Profile{}
		newProfile.Id = userID

		err = db.SetUserProfileWithByteData(c.BodyRaw(), userID)
		if err != nil {
			fmt.Println("Error Setting the profile: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Profile is set successfully!")
	}
}
func GetAdditionalUserInfo() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}

		additional_profile, err := db.GetAdditionalInfoProfileByID(userID)
		if err != nil {
			fmt.Println("error while fetching additional profile: ", err)
			return c.Status(500).SendString("unable to fetch user additional profile!")
		}

		return c.JSON(additional_profile)
	}
}
func SetAdditionalUserInfo() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}
		err = db.SetAdditionalUserProfileWithByteData(c.BodyRaw(), userID)
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}
		return c.SendString("User Additional profile updated Successfully!")
	}
}
func GetUserSettings() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}

		additional_profile, err := db.GetUserSettingsByID(userID)
		if err != nil {
			fmt.Println("error while fetching additional profile: ", err)
			return c.Status(500).SendString("unable to fetch user additional profile!")
		}

		return c.JSON(additional_profile)
	}
}
func SetUserSettings() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}
		err = db.SetUserSettingsWithByteData(c.BodyRaw(), userID)
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}
		return c.SendString("User Additional profile updated Successfully!")
	}
}
func GetUserData() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}

		additional_profile, err := db.GetUserDataByID(userID)
		if err != nil {
			fmt.Println("error while fetching additional profile: ", err)
			return c.Status(500).SendString("unable to fetch user additional profile!")
		}

		return c.JSON(additional_profile)
	}
}
func SetUserData() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}
		err = db.SetUserDataWithByteData(c.BodyRaw(), userID)
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}
		return c.SendString("User Additional profile updated Successfully!")
	}
}
