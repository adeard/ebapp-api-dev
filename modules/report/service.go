package report

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	GetByPlant(ids []string) ([]domain.ListReport, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetByPlant(ids []string) ([]domain.ListReport, error) {
	projects, err := s.repository.FindByPlant(ids)
	return projects, err
}
