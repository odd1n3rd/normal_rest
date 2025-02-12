package repository

import (
	"github.com/odd1n3rd/default_web/internal/entity"
	"gorm.io/gorm"
)

type OrderRepos interface {
	Create(order *entity.Order) error
	GetByID(id uint) (*entity.Order, error)
	GetAll() ([]entity.Order, error)
	Update(order *entity.Order) error
	DeleteByID(id uint) error
}

type orderRepos struct {
	db *gorm.DB
}

func NewOrderRepos(db *gorm.DB) OrderRepos {
	return &orderRepos{db: db}
}

func (r *orderRepos) Create(order *entity.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepos) GetByID(id uint) (*entity.Order, error) {
	var order entity.Order
	err := r.db.Preload("Users").First(&order, id).Error
	return &order, err
}

func (r *orderRepos) Update(order *entity.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepos) DeleteByID(id uint) error {
	return r.db.Delete(&entity.Order{}, id).Error
}

func (r *orderRepos) GetAll() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Find(&orders).Error
	return orders, err
}
