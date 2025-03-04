package service

import (
	"github.com/odd1n3rd/default_web/internal/entity"
	"github.com/odd1n3rd/default_web/internal/repository"
)

type OrderService interface {
	GetAllOrders() ([]entity.Order, error)
	GetOrderByID(id uint) (*entity.Order, error)
	CreateOrder(order *entity.Order) error
	UpdateOrderInfo(order *entity.Order, id int) error
	DeleteOrder(id uint) error
	GetAllUserOrders(id uint) ([]entity.Order, error)
}

type orderService struct {
	repos repository.OrderRepos
}

func NewOrderService(r repository.OrderRepos) OrderService {
	return &orderService{repos: r}
}

func (s *orderService) GetAllOrders() ([]entity.Order, error) {
	return s.repos.GetAll()
}

func (s *orderService) GetOrderByID(id uint) (*entity.Order, error) {
	return s.repos.GetByID(id)
}

func (s *orderService) GetAllUserOrders(id uint) ([]entity.Order, error) {
	return s.repos.GetAllByUserId(id)
}

func (s *orderService) CreateOrder(order *entity.Order) error {
	return s.repos.Create(order)
}

func (s *orderService) UpdateOrderInfo(order *entity.Order, id int) error {
	return s.repos.UpdateByID(order, id)
}

func (s *orderService) DeleteOrder(id uint) error {
	return s.repos.DeleteByID(id)
}
