package services

import "fintrack-backend/models"

type AuthService interface {
	SignUp(username, email, password string) (*models.User, error)
	Login(email, password string) (string, error)
}
