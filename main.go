package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(mysql.Open("root:@/go_auth_react"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database!")
	}

	app := fiber.New()

	app.Get("/", home)

	app.Listen(":8000")
}

func home(c *fiber.Ctx) error {
	return c.SendString("Hello world ~")
}
