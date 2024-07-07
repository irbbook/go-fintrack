package repositories

import (
	"fintrack-backend/models"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

// FindUserByEmail implements AuthRepository.
func (a *authRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := a.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *authRepository) CreateUser(user *models.User) error {
	// CreateUser implements AuthRepository.
	return a.db.Create(user).Error
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}
