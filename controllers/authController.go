package controllers

import (
	"go-auth-react/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user models.User

	user.FirstName = "John"
	user.LastName = "Doe"
	user.Email = "john@mail.com"
	user.Password = "pass"

	return c.JSON(user)
}
