package main

import (
	"fintrack-backend/config"
	"fintrack-backend/controllers"
	"fintrack-backend/repositories"
	"fintrack-backend/routes"
	"fintrack-backend/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	app := fiber.New()

	config.InitDatabase()

	authRepo := repositories.NewAuthRepository(config.DB)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	routes.Setup(app, authController)

	log.Fatal(app.Listen(":3000"))
}
