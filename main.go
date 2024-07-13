package main

import (
	"fintrack-backend/config"
	"fintrack-backend/controllers"
	"fintrack-backend/repositories"
	"fintrack-backend/routes"
	"fintrack-backend/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	app := fiber.New()

	config.InitDatabase()
	config.InitRedis()

	authRepo := repositories.NewAuthRepository(config.DB)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	routes.Setup(app, authController)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
