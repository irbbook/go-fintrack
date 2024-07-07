package routes

import (
	"fintrack-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, authController *controllers.AuthController) {
	app.Post("/signup", authController.SignUp)
	app.Post("/login", authController.Login)
}
