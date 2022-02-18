package seeders

import (
	"github.com/furqonzt99/refresh-token/constants"
	"github.com/furqonzt99/refresh-token/models"
	"github.com/furqonzt99/refresh-token/services"
	"gorm.io/gorm"
)

func AdminSeeder(db *gorm.DB) {
	password, _ := services.Hashpwd(constants.ADMIN_PASSWORD)

	admin := models.User{
		Name:     "Admin",
		Email:    constants.ADMIN_EMAIL,
		Password: password,
		Role:     "admin",
	}

	db.Create(&admin)
}
