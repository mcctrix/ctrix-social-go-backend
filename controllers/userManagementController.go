package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	db "github.com/mcctrix/ctrix-social-go-backend/db/v1"
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
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch user profile!")
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

		err = db.SetUserProfileWithByteData(c.BodyRaw(), userID)
		if err != nil {
			fmt.Println("Error Setting the profile: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Profile is updated successfully!")
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

func UpdateAdditionalUserInfo() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}
		err = db.UpdateAdditionalUserProfileWithByteData(c.BodyRaw(), userID)
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
			return c.Status(500).SendString("unable to fetch user settings!")
		}

		return c.JSON(additional_profile)
	}
}

func UpdateUserSettings() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}
		err = db.UpdateUserSettingsWithByteData(c.BodyRaw(), userID)
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
			return c.Status(500).SendString("unable to fetch user data!")
		}

		return c.JSON(additional_profile)
	}
}

func UpdateUserData() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := utils.GetUserIDWithToken(c.Cookies("auth_token"))
		if err != nil {
			fmt.Println("unable to fetch user with this Token: ", err)
			return c.Status(401).SendString("unable to fetch user with this Token!")
		}
		err = db.UpdateUserDataWithByteData(c.BodyRaw(), userID)
		if err != nil {
			fmt.Println("Error Setting the user data: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("User Data updated Successfully!")
	}
}
