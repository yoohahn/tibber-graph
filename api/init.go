package api

import (
	"log"
	"server/tibber"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func securityHeaders() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
		c.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}

		return c.Next()
	}
}
func registerApp() {
	app := fiber.New(fiber.Config{GETOnly: true, CompressedFileSuffix: ".gz"})
	app.Use(securityHeaders())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		result, _ := tibber.GetPrice()
		return c.SendString(result)
	})

	log.Fatal(app.Listen(":4321"))
}

func InitApi() {
	registerApp()
}
