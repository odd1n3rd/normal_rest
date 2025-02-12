package repository

import (
	"github.com/odd1n3rd/default_web/internal/entity"
	"gorm.io/gorm"
)

type UserRepos interface {
	Create(user *entity.User) error
	GetByID(id uint) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Update(user *entity.User) error
	DeleteByID(id uint) error
}

type userRepos struct {
	db *gorm.DB
}

func NewUserRepos(db *gorm.DB) UserRepos {
	return &userRepos{db: db}
}

func (r *userRepos) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepos) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.Preload("Orders").First(&user, id).Error
	return &user, err
}

func (r *userRepos) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepos) DeleteByID(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}

func (r *userRepos) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}
