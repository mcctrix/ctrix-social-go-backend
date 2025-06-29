package auth

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/models"
	"github.com/mcctrix/ctrix-social-go-backend/internal/domain/services"
	db "github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/auth"
	"gorm.io/gorm"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthService(userService *services.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	// If auth token exist we don't let login again
	authToken := c.Cookies("auth_token")
	if authToken != "" {
		return c.SendString("User already logged in!")
	}

	username := strings.ToLower(c.FormValue("username"))
	password := c.FormValue("password")

	user, err := h.userService.AuthenticateUser(username, password)
	if err != nil {
		fmt.Println("Error while authenticating user: ", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	gnToken, err := h.userService.GenerateJwtToken(user)
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

func (h *AuthHandler) SignUp(c fiber.Ctx) error {

	authToken := c.Cookies("auth_token")
	if authToken != "" {
		return c.SendString("User already logged in!")
	}

	email := strings.ToLower(c.FormValue("email"))
	username := strings.ToLower(c.FormValue("username"))
	password := c.FormValue("password")

	user, err := h.userService.RegisterUser(email, username, password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	gnToken, err := h.userService.GenerateJwtToken(user)
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

func (h *AuthHandler) Logout(c fiber.Ctx) error {
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

func (h *AuthHandler) RefreshToken(c fiber.Ctx) error {
	auth_token := c.Cookies("auth_token")

	if auth_token == "" {
		return c.SendString("No token found!")
	}

	jwtToken, err := auth.GetJwtToken(auth_token)
	if err != nil {
		fmt.Println("Error while converting string to token: ", err)
		return fiber.ErrInternalServerError
	}
	if !jwtToken.Valid {
		return c.SendString("Invalid Token!")
	}

	userID := auth.GetClaimData(jwtToken, "aud")

	if userID == "" {
		return fiber.ErrInternalServerError
	}

	db := db.GetDB()

	var user *models.User = &models.User{}

	if err = db.Table("user_auth").Where("id = ?", userID).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendString("User not Found!")
		}
		return fiber.ErrInternalServerError
	}

	gnToken, err := auth.GenerateJwtToken(user)
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

func (h *AuthHandler) ForgetPassword(c fiber.Ctx) error {
	email := c.FormValue("email")

	db := db.GetDB()

	var user *models.User = &models.User{}
	if err := db.Table("user_auth").Where(struct{ Email string }{Email: email}).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendString("User not found with this email!")
		}
	}

	return c.SendString("User exist but i can't forget password for you!")
}

func (h *AuthHandler) ResetPassword(c fiber.Ctx) error {
	email := c.FormValue("email")
	oldPassword := c.FormValue("old_password")
	newPassword := c.FormValue("new_password")

	db := db.GetDB()

	var user *models.User = &models.User{}
	if err := db.Table("user_auth").Where(struct {
		Email    string
		Password string
	}{Email: email, Password: oldPassword}).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendString("User not found with this password!")
		}
	}

	user.Password = newPassword
	if err := db.Table("user_auth").Save(user).Error; err != nil {
		fmt.Println("Error while updating password: ", err)
		return fiber.ErrInternalServerError
	}

	return c.SendString("Password resetted!")
}
