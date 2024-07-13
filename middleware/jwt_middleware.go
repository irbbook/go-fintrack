package middleware

import (
	"context"
	"fintrack-backend/config"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or malformed JWT"})
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired JWT"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check if the token is blacklisted
			if isTokenBlacklisted(tokenString) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token has been logged out"})
			}

			// Pass user_id to the next handler
			c.Locals("user_id", claims["user_id"])
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired JWT"})
	}
}

func isTokenBlacklisted(tokenString string) bool {
	_, err := config.RedisClient.Get(context.Background(), tokenString).Result()
	return err == nil // If no error, the token is blacklisted
}
