package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/r-52/aurora/aurora_http"
	"github.com/r-52/aurora/database"
)

func main() {
	app := fiber.New()
	aurora_http.CreateSession()

	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
