package seeders

import (
	"github.com/furqonzt99/refresh-token/helpers"
	"github.com/furqonzt99/refresh-token/models"
	"gorm.io/gorm"
)

func AdminSeeder(db *gorm.DB) {
	password, _ := helpers.Hashpwd("1234qwer")

	admin := models.User{
		Name:     "Admin",
		Email:    "admin@admin.com",
		Password: password,
		Role:     "admin",
	}

	db.Create(&admin)
}
