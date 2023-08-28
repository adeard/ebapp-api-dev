package user

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	GetAll(input domain.UserRequest) ([]domain.User, error)
	GetByID(userId string) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.UserRequest) ([]domain.User, error) {
	users, err := s.repository.FindAll(input)
	return users, err
}

func (s *service) GetByID(userId string) (domain.User, error) {
	users, err := s.repository.FindByUserId(userId)
	return users, err
}
