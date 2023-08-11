package listproject

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	GetAll(input domain.ListProjectRequest) ([]domain.ListProject, error)
	GetByID(id string) (domain.ListProject, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.ListProjectRequest) ([]domain.ListProject, error) {
	listProjects, err := s.repository.FindAll(input)
	return listProjects, err
}

func (s *service) GetByID(id string) (domain.ListProject, error) {
	project, err := s.repository.FindById(id)
	return project, err
}
