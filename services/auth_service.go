package services

import (
	"errors"
	"fintrack-backend/models"
	"fintrack-backend/repositories"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo repositories.AuthRepository
}

// Login implements AuthService.
func (a *authService) Login(email string, password string) (string, error) {
	user, err := a.repo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := generateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

// SignUp implements AuthService.
func (a *authService) SignUp(username string, email string, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	if err := a.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo}
}

func generateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
