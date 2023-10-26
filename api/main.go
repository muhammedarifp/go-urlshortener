package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/muhammedarifp/go-urlshortner/api/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1", routes.ShortenUrl)
}

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("err is : ", err.Error())
	// }

	app := fiber.New()

	setupRoutes(app)
	app.Use(logger.New())

	app.Listen(":3000")
}
