package seeders

import (
	"github.com/furqonzt99/refresh-token/models"
	"github.com/furqonzt99/refresh-token/services"
	"gorm.io/gorm"
)

func AdminSeeder(db *gorm.DB) {
	password, _ := services.Hashpwd("1234qwer")

	admin := models.User{
		Name:     "Admin",
		Email:    "admin@admin.com",
		Password: password,
		Role:     "admin",
	}

	db.Create(&admin)
}
