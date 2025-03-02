package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func Login() fiber.Handler {
	return func(c fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		fmt.Println(username)
		fmt.Println(password)
		return c.SendString("User: " + username)
	}
}

func SignUp() fiber.Handler {
	return func(c fiber.Ctx) error {
		email := c.FormValue("email")
		username := c.FormValue("username")
		password := c.FormValue("password")
		fmt.Println(email)
		fmt.Println(username)
		fmt.Println(password)
		return c.SendString("User: " + username + " Found!")
	}
}
func Logout() fiber.Handler {
	return func(c fiber.Ctx) error {

		return nil
	}
}
