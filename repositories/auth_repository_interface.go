package repositories

import "fintrack-backend/models"

type AuthRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	GetUsers() ([]models.User, error)
}
