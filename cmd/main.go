package main

import (
	"github.com/betbil/go-docker/database"
	"github.com/betbil/go-docker/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.HomeHandler)
	app.Post("/facts", handlers.CreateFactHandler)
	app.Get("/list", handlers.ListFactHandler)
}
