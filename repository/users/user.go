package users

import (
	"github.com/furqonzt99/refresh-token/models"
	"gorm.io/gorm"
)

type UserInterface interface {
	Create(user models.User) (models.User, error)
	ReadAll() ([]models.User, error)
	ReadOne(id string) (models.User, error)
	Update(id string, updateUser models.User) (models.User, error)
	Delete(id string) (models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user models.User) (models.User, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) ReadAll() ([]models.User, error) {
	var users []models.User

	ur.db.Find(&users)

	return users, nil
}

func (ur *UserRepository) ReadOne(id string) (models.User, error) {
	var user models.User

	if err := ur.db.First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Update(id string, updateUser models.User) (models.User, error) {
	var user models.User

	if err := ur.db.First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}

	ur.db.Model(&user).Updates(updateUser)

	return updateUser, nil
}

func (ur *UserRepository) Delete(id string) (models.User, error) {
	var user models.User

	if err := ur.db.First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}

	ur.db.Delete(&user)

	return user, nil
}
