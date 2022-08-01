package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/r-52/aurora/aurora_http"
	"github.com/r-52/aurora/components/configuration"
	"github.com/r-52/aurora/database"
)

func main() {
	app := fiber.New()

	prod, err := os.ReadFile("./config.prod.ini")
	if err != nil {
		panic("could not read prod ini")
	}
	dev, err := os.ReadFile("./config.dev.ini")
	if err != nil {
		panic("could not read dev ini")
	}
	configuration.BuildIni(prod, dev)

	aurora_http.CreateSession()

	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
