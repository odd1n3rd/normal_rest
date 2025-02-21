package entity

import (
	"time"
	// "gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID        uint `gorm:"primaryKey" gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt ``

	Name   string  `json:"name"`
	Email  string  `json:"email" gorm:"unique"`
	Orders []Order `json:"orders"`
}
