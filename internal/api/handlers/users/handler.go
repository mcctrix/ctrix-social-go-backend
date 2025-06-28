package users

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	repo "github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database/repositories"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
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

		err = repo.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_additional_info")
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}

		err = repo.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_profile")
		if err != nil {
			fmt.Println("Error Setting the user profile: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Profile Setup Successfully!")
	}
}

func GetCurrentUserProfile() fiber.Handler {
	return func(c fiber.Ctx) error {

		profile, err := repo.GetUserData(c.Locals("userID").(string), "user_profile", []string{"first_name", "last_name", "avatar", "last_seen", "verified_user"})
		if err != nil {
			fmt.Println("unable to fetch profile: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch user profile!")
		}
		userAuth, err := repo.GetUserData(c.Locals("userID").(string), "user_auth", []string{"email", "username", "created_at"})
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
		profile, err := repo.GetUserData(userID, "user_profile", []string{"first_name", "last_name", "avatar", "last_seen"})
		if err != nil {
			fmt.Println("unable to fetch profile: ", err)
			return c.Status(fiber.StatusNotFound).SendString("unable to fetch user profile!")
		}
		userAuth, err := repo.GetUserData(userID, "user_auth", []string{"email", "username"})
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
			Avatar          string `json:"avatar,omitempty"`
		}{}

		rawForm, err := utils.ClearStruct(dataInterface, c.BodyRaw())
		if err != nil {
			fmt.Println("Error Marshalling the data: ", err)
			return fiber.ErrInternalServerError
		}

		err = repo.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_profile")
		if err != nil {
			fmt.Println("Error Setting the profile: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Profile is updated successfully!")
	}
}

func GetAdditionalUserInfo() fiber.Handler {
	return func(c fiber.Ctx) error {
		fmt.Println("Get Additional User Info")
		additional_profile, err := repo.GetUserData(c.Locals("userID").(string), "user_additional_info", []string{"hobbies", "relation_status", "dob", "bio", "gender"})
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

		err = repo.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_additional_info")
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}
		return c.SendString("User Additional profile updated Successfully!")
	}
}

func GetUserSettings() fiber.Handler {
	return func(c fiber.Ctx) error {
		additional_profile, err := repo.GetUserData(c.Locals("userID").(string), "user_settings", []string{"hide_post", "hide_story", "block_user", "show_online"})
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

		err = repo.UpdateTableWithByteData(rawForm, c.Locals("userID").(string), "user_settings")
		if err != nil {
			fmt.Println("Error Setting the additional profile: ", err)
			return fiber.ErrInternalServerError
		}
		return c.SendString("User Additional profile updated Successfully!")
	}
}

func FollowUser() fiber.Handler {
	return func(c fiber.Ctx) error {
		follow_id := c.Params("userid")
		following_id := c.Locals("userID").(string)

		err := repo.FollowUser(follow_id, following_id)
		if err != nil {
			fmt.Println("Error while following user: ", err)
			return err
		}

		return c.SendString("User Followed Successfully!")
	}
}

func UnfollowUser() fiber.Handler {
	return func(c fiber.Ctx) error {
		follow_id := c.Params("userid")
		following_id := c.Locals("userID").(string)

		err := repo.UnfollowUser(follow_id, following_id)
		if err != nil {
			fmt.Println("Error while unfollowing user: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("User Unfollowed Successfully!")
	}
}

func CheckFollowing() fiber.Handler {
	return func(c fiber.Ctx) error {
		follow_id := c.Params("userid")
		following_id := c.Locals("userID").(string)

		follow, err := repo.CheckFollowing(follow_id, following_id)
		if err != nil {
			fmt.Println("Error while checking following: ", err)
			return c.Status(404).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(follow)
	}
}

func GetFollowAndFollowing() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Params("userID")

		follow, err := repo.GetFollowAndFollowing(userID)
		if err != nil {
			fmt.Println("Error while getting follow and following: ", err)
			return c.Status(404).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(follow)
	}
}