package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	// gorm.Model
	ID        uint `gorm:"primaryKey" gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	ProductName string `json:"product_name"`
	Amount      int    `json:"amount"`
	Cost        int    `json:"cost"`
	UserID      int    `json:"user_id"`
}
