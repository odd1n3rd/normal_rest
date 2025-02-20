package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductName string  `json:"product_name"`
	Amount      int     `json:"amount"`
	Cost        float32 `json:"cost"`
	UserID      int     `json:"user_id"`
}
