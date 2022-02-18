package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Email     string
	Password  string
	Role      string `gorm:"default:user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now().Local()
	return nil
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	u.UpdatedAt = time.Now().Local()
	return nil
}
