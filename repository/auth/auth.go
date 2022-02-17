package auth

import (
	"github.com/furqonzt99/refresh-token/models"
	"gorm.io/gorm"
)

type AuthInterface interface {
	Login(email string) (models.User, error)
	Register(user models.User) (models.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) Login(email string) (models.User, error) {
	var user models.User

	if err := ar.db.First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ar *AuthRepository) Register(user models.User) (models.User, error) {

	if err := ar.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
