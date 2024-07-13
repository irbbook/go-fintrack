package routes

import (
	"fintrack-backend/controllers"
	"fintrack-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, authController *controllers.AuthController) {
	app.Post("/signup", authController.SignUp)
	app.Post("/login", authController.Login)
	app.Get("/users", middleware.JWTMiddleware(), authController.GetUsers)
}
