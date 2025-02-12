package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string  `json:"name"`
	Email  int     `json"email"`
	Orders []Order `json:"orders"`
}
