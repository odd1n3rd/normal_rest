package service

import (
	"github.com/odd1n3rd/default_web/internal/entity"
	"github.com/odd1n3rd/default_web/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUserInfo(user *entity.User, id int) error
	DeleteUser(id uint) error
}

type userService struct {
	repos repository.UserRepos
}

func NewUserService(r repository.UserRepos) UserService {
	return &userService{repos: r}
}

func (s *userService) GetAllUsers() ([]entity.User, error) {
	return s.repos.GetAll()
}

func (s *userService) GetUserByID(id uint) (*entity.User, error) {
	return s.repos.GetByID(id)
}

func (s *userService) CreateUser(user *entity.User) error {
	return s.repos.Create(user)
}

func (s *userService) UpdateUserInfo(user *entity.User, id int) error {
	return s.repos.UpdateByID(user, id)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repos.DeleteByID(id)
}
