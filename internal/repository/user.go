package repository

import (
	"github.com/odd1n3rd/default_web/internal/entity"
	"gorm.io/gorm"
)

type UserRepos interface {
	Create(user *entity.User) error
	GetByID(id uint) (*entity.User, error)
	GetAll() ([]entity.User, error)
	UpdateByID(user *entity.User, ID int) error
	DeleteByID(id uint) error
}

type userRepos struct {
	db *gorm.DB
}

func NewUserRepos(db *gorm.DB) UserRepos {
	return &userRepos{db: db}
}

func (r *userRepos) Create(user *entity.User) error {
	return r.db.Preload("orders").Create(user).Error
}

func (r *userRepos) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.Preload("Orders").First(&user, id).Error
	return &user, err
}

func (r *userRepos) UpdateByID(user *entity.User, ID int) error {
	// var updatedUser entity.User
	// if err := r.db.First(&updatedUser, ID); err != nil {
	// 	return err.Error
	// }

	// updatedUser.Email = user.Email
	// updatedUser.Name = user.Name
	if err := r.db.Model(&user).Where("id=?", ID).Updates(map[string]any{
		"Name":  user.Name,
		"Email": user.Email,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepos) DeleteByID(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}

func (r *userRepos) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Preload("Orders").Find(&users).Error
	return users, err
}
