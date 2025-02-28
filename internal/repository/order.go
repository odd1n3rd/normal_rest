package repository

import (
	"github.com/odd1n3rd/default_web/internal/entity"
	"gorm.io/gorm"
)

type OrderRepos interface {
	Create(order *entity.Order) error
	GetByID(id uint) (*entity.Order, error)
	GetAll() ([]entity.Order, error)
	UpdateByID(order *entity.Order, id int) error
	DeleteByID(id uint) error
	GetAllByUserId(id uint) ([]entity.Order, error)
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

func (r *orderRepos) UpdateByID(order *entity.Order, id int) error {
	// var updatedOrd entity.Order
	// r.db.First(&updatedOrd, id)
	// updatedOrd.ProductName = order.ProductName
	// updatedOrd.Amount = order.Amount
	// updatedOrd.Cost = order.Cost
	return r.db.Model(&entity.Order{}).Where("id = ?", id).Save(&entity.Order{ProductName: order.ProductName,
		Cost:   order.Cost,
		Amount: order.Amount,
		UserID: order.UserID,
	}).Error
}

func (r *orderRepos) DeleteByID(id uint) error {
	return r.db.Delete(&entity.Order{}, id).Error
}

func (r *orderRepos) GetAll() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *orderRepos) GetAllByUserId(id uint) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Where("user_id = ?", id).Find(&orders).Error
	return orders, err
}
