package controllers

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	db "github.com/mcctrix/ctrix-social-go-backend/db/v1"
	"github.com/mcctrix/ctrix-social-go-backend/models"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
	"gorm.io/gorm"
)

func SignUp() fiber.Handler {
	return func(c fiber.Ctx) error {

		// If auth token exist we don't let login again
		authToken := c.Cookies("auth_token")
		if authToken != "" {
			// return c.Redirect().To("/")
			return c.SendString("User already logged in!")
		}

		email := strings.ToLower(c.FormValue("email"))
		username := strings.ToLower(c.FormValue("username"))
		password := c.FormValue("password")

		user := &models.User_Auth{}
		user.Email = email
		user.Username = username
		user.Password = password
		user.Id = uuid.New().String()
		user.Created_at = time.Now()

		dbInstance, err := db.DBConnection()
		if err != nil {
			fmt.Println(err)
			return fiber.ErrInternalServerError
		}
		if err = dbInstance.Table("user_auth").Create(user).Error; err != nil {
			fmt.Println("Error in Creating New User:", err.Error())
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return c.Status(400).SendString("Duplicate Email or Username.")
			}

			return fiber.ErrInternalServerError
		}

		gnToken, err := utils.GenerateJwtToken(user)
		if err != nil {
			fmt.Println("Error While Generating token: ", err)
			return fiber.ErrInternalServerError
		}

		expireTime := time.Unix(gnToken.Exp_Time, 0)

		c.Cookie(&fiber.Cookie{
			Name:     "auth_token",
			Value:    gnToken.StringToken,
			Path:     "/",
			HTTPOnly: true,
			Secure:   true,
			Domain: func() string {
				if os.Getenv("APP_ENV") == "dev" {
					return ""
				}
				return "ctrix-social.vercel.app"
			}(),
			SameSite: "None",
			Expires:  expireTime,
		})

		go func() {
			err := db.InitNewUser(user.Id)
			if err != nil {
				fmt.Println("Error while initializing new user: ", err)
			}
		}()

		return c.JSON(fiber.Map{"tokenValue": gnToken.StringToken, "Expires": expireTime})
	}
}

func Login() fiber.Handler {
	return func(c fiber.Ctx) error {
		// If auth token exist we don't let login again
		authToken := c.Cookies("auth_token")
		if authToken != "" {
			// return c.Redirect().To("/")
			return c.SendString("User already logged in!")
		}

		dbConn, err := db.DBConnection()
		if err != nil {
			fmt.Println(err)
			return fiber.ErrInternalServerError
		}
		username := strings.ToLower(c.FormValue("username"))
		password := c.FormValue("password")

		user := &models.User_Auth{}

		whereConditionData := &struct {
			Username string
			// Password string
		}{
			Username: username,
			// Password: password,
		}

		// select * from user_auth where password="12345678"
		if err = dbConn.Table("user_auth").Where(whereConditionData).First(user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).SendString("User not Found!")
			}
			// return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		if password != user.Password {
			return c.Status(fiber.StatusUnauthorized).SendString("Incorrect Password for " + user.Username)
		}
		// Prank for sending passowrd of a user
		// if user.Username != username {
		// 	return c.SendString("Password is of user: " + user.Username)
		// }

		gnToken, err := utils.GenerateJwtToken(user)
		if err != nil {
			fmt.Println("Error While Generating token: ", err)
			return fiber.ErrInternalServerError
		}

		expireTime := time.Unix(gnToken.Exp_Time, 0)

		c.Cookie(&fiber.Cookie{
			Name:     "auth_token",
			Value:    gnToken.StringToken,
			Path:     "/",
			HTTPOnly: true,
			Secure:   false,
			Domain: func() string {
				if os.Getenv("APP_ENV") == "dev" {
					return ""
				}
				return "ctrix-social.vercel.app"
			}(),
			SameSite: "Lax",
			Expires:  expireTime,
		})

		return c.JSON(fiber.Map{"tokenValue": gnToken.StringToken, "Expires": expireTime})
	}
}
func Logout() fiber.Handler {
	return func(c fiber.Ctx) error {
		c.ClearCookie("auth_token")

		c.Cookie(&fiber.Cookie{
			Name: "auth_token",
			// Set expiry date to the past
			Expires:  time.Now().Add(-(time.Hour * 2)),
			HTTPOnly: true,
			SameSite: "lax",
		})
		return c.SendString("User logged out Successfully!")
	}
}

func RefreshToken() fiber.Handler {
	return func(c fiber.Ctx) error {
		auth_token := c.Cookies("auth_token")

		if auth_token == "" {
			return c.SendString("No token found!")
		}

		jwtToken, err := utils.GetJwtToken(auth_token)
		if err != nil {
			fmt.Println("Error while converting string to token: ", err)
			return fiber.ErrInternalServerError
		}
		if !jwtToken.Valid {
			return c.SendString("Invalid Token!")
		}

		userID := utils.GetClaimData(jwtToken, "aud")

		if userID == "" {
			return fiber.ErrInternalServerError
		}

		db, err := db.DBConnection()
		if err != nil {
			fmt.Println(err)
			return fiber.ErrInternalServerError
		}
		var user *models.User_Auth = &models.User_Auth{}
		if err = db.Table("user_auth").Where(struct{ ID string }{ID: userID}).First(user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.SendString("user not found for refreshing token, is that even possible lol!")
			}
		}
		gnToken, err := utils.GenerateJwtToken(user)
		if err != nil {
			fmt.Println("Error While Generating token from refresh token: ", err)
			return fiber.ErrInternalServerError
		}
		c.Cookie(&fiber.Cookie{
			Name:     "auth_token",
			Value:    gnToken.StringToken,
			Path:     "/",
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Lax",
			Expires:  time.Unix(gnToken.Exp_Time, 0),
		})
		return c.SendString("Token is Refreshed!")

	}
}
func ForgetPassword() fiber.Handler {
	return func(c fiber.Ctx) error {
		email := c.FormValue("email")

		db, err := db.DBConnection()
		if err != nil {
			fmt.Println("Error while connecting to db: ", err)
			return fiber.ErrInternalServerError
		}
		var user *models.User_Auth = &models.User_Auth{}
		if err = db.Table("user_auth").Where(struct{ Email string }{Email: email}).First(user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.SendString("User not found with this email!")
			}
		}

		return c.SendString("User exist but i can't forget password for you!")
	}
}
func ResetPassword() fiber.Handler {
	return func(c fiber.Ctx) error {
		email := c.FormValue("email")
		oldPassword := c.FormValue("old_password")
		newPassword := c.FormValue("new_password")

		db, err := db.DBConnection()
		if err != nil {
			fmt.Println("Error while connecting to db: ", err)
			return fiber.ErrInternalServerError
		}
		var user *models.User_Auth = &models.User_Auth{}
		if err = db.Table("user_auth").Where(struct {
			Email    string
			Password string
		}{Email: email, Password: oldPassword}).First(user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.SendString("User not found with this password!")
			}
		}

		user.Password = newPassword
		if err = db.Table("user_auth").Save(user).Error; err != nil {
			fmt.Println("Error while updating password: ", err)
			return fiber.ErrInternalServerError
		}

		return c.SendString("Password resetted!")
	}
}
