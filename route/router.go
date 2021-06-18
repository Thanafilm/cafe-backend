package route

import (
	"cafe-backend/controller"
	"cafe-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	app.Post("/signup", controller.Signup)
	app.Post("/login", controller.Login)

	app.Use(middleware.Authenticated)

	app.Get("/user", controller.Users)
	app.Get("/logout", controller.Logout)
	app.Get("/allUser", controller.AllUser)
}
