package repository

import (
	"github.com/roshith/uber-fx/internals/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindAll() ([]models.User, error) {
	var users []models.User
	return users, r.db.Find(&users).Error
}
