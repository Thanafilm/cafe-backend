package main

import (
	"cafe-backend/connection"
	"cafe-backend/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	route.Router(app)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	connection.ConnectDB()
	app.Listen(":3000")
}
