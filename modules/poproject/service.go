package poproject

import "ebapp-api-dev/domain"

type Service interface {
	GetAll(input domain.PoProjectRequest) ([]domain.PoProject, error)
	GetByPo(po string) ([]domain.PoProject, error)
	Store(input domain.PoProject) (domain.PoProject, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll(input domain.PoProjectRequest) ([]domain.PoProject, error) {
	poProject, err := s.repository.FindAll(input)
	return poProject, err
}

func (s *service) GetByPo(po string) ([]domain.PoProject, error) {
	poProject, err := s.repository.FindByPo(po)
	return poProject, err
}

func (s *service) Store(input domain.PoProject) (domain.PoProject, error) {
	poProject, err := s.repository.Store(input)
	return poProject, err
}
