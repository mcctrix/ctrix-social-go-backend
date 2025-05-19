package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	db "github.com/mcctrix/ctrix-social-go-backend/db/v1"
)

func GetCurrentUserProfile() fiber.Handler {
	return func(c fiber.Ctx) error {
		profile, err := db.GetUserProfileByID(c.Locals("userID").(string))
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

		err := db.SetUserProfileWithByteData(c.BodyRaw(), c.Locals("userID").(string))
		if err != nil {
			fmt.Println("Error Setting the profile: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Profile is updated successfully!")
	}
}

func GetAdditionalUserInfo() fiber.Handler {
	return func(c fiber.Ctx) error {

		additional_profile, err := db.GetAdditionalInfoProfileByID(c.Locals("userID").(string))
		if err != nil {
			fmt.Println("error while fetching additional profile: ", err)
			return c.Status(500).SendString("unable to fetch user additional profile!")
		}

		return c.JSON(additional_profile)
	}
}

func UpdateAdditionalUserInfo() fiber.Handler {
	return func(c fiber.Ctx) error {
		err := db.UpdateAdditionalUserProfileWithByteData(c.BodyRaw(), c.Locals("userID").(string))
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}
		return c.SendString("User Additional profile updated Successfully!")
	}
}

func GetUserSettings() fiber.Handler {
	return func(c fiber.Ctx) error {
		additional_profile, err := db.GetUserSettingsByID(c.Locals("userID").(string))
		if err != nil {
			fmt.Println("error while fetching additional profile: ", err)
			return c.Status(500).SendString("unable to fetch user settings!")
		}

		return c.JSON(additional_profile)
	}
}

func UpdateUserSettings() fiber.Handler {
	return func(c fiber.Ctx) error {
		err := db.UpdateUserSettingsWithByteData(c.BodyRaw(), c.Locals("userID").(string))
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}
		return c.SendString("User Additional profile updated Successfully!")
	}
}

func GetUserData() fiber.Handler {
	return func(c fiber.Ctx) error {

		additional_profile, err := db.GetUserDataByID(c.Locals("userID").(string))
		if err != nil {
			fmt.Println("error while fetching additional profile: ", err)
			return c.Status(500).SendString("unable to fetch user data!")
		}

		return c.JSON(additional_profile)
	}
}

func UpdateUserData() fiber.Handler {
	return func(c fiber.Ctx) error {

		err := db.UpdateUserDataWithByteData(c.BodyRaw(), c.Locals("userID").(string))
		if err != nil {
			fmt.Println("Error Setting the user data: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("User Data updated Successfully!")
	}
}
