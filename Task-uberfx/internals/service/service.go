package service

import (
	"github.com/roshith/uber-fx/internals/models"
	"github.com/roshith/uber-fx/internals/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (s *UserService) Create(name, email string) (*models.User, error) {
	user := &models.User{
		Name:  name,
		Email: email,
	}

	err := s.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FindAll() ([]models.User, error) {
	return s.repository.FindAll()
}
