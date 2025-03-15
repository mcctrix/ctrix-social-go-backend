package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mcctrix/ctrix-social-go-backend/db"
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

		user := &models.User{}
		user.Email = c.FormValue("email")
		user.Username = c.FormValue("username")
		user.Password = c.FormValue("password")
		user.ID = uuid.New().String()
		user.Created_at = time.Now()

		db, err := db.DBConnection()
		if err != nil {
			return fiber.ErrInternalServerError
		}
		if err = db.Table("user_auth").Create(user).Error; err != nil {
			fmt.Println("Error in Creating New User:", err.Error())
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return c.Status(400).JSON(map[string]string{
					"error": "Duplicate Email or Username.",
				})
			}

			return fiber.ErrInternalServerError
		}

		gnToken, err := utils.GenerateJwtToken(user)
		if err != nil {
			fmt.Println("Error While Generating token: ", err)
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

		return c.SendString("User Created Successfully!")
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
			return fiber.ErrInternalServerError
		}
		username := c.FormValue("username")
		password := c.FormValue("password")

		user := &models.User{}

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
				return c.SendString("User not Found!")
			}
		}
		if password != user.Password {
			return c.SendString("Incorrect Password for " + user.Username)
		}
		// Prank for sending passowrd of a user
		// if user.Username != username {
		// 	return c.SendString("Password is of user: " + user.Username)
		// }

		gnToken, err := utils.GenerateJwtToken(user)

		if err != nil {
			fmt.Println("Error Creating Jwt while login: ", err)
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

		return c.SendString("User logged in Succesfully!")
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

		if claim, ok := jwtToken.Claims.(jwt.MapClaims); ok {
			userID := claim["aud"].(string)
			db, err := db.DBConnection()
			if err != nil {
				fmt.Println(err)
				return fiber.ErrInternalServerError
			}
			var user *models.User = &models.User{}
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

		} else {
			return fiber.ErrInternalServerError
		}
	}
}
