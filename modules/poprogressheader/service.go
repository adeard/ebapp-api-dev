package poprogressheader

import "ebapp-api-dev/domain"

type Service interface {
	FindProg(id string) (domain.PoProgressHeader, error)
	FindAllProg(id string) ([]domain.PoProgressHeader, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) FindProg(id string) (domain.PoProgressHeader, error) {
	data, err := s.repository.FindProg(id)
	return data, err
}

func (s *service) FindAllProg(id string) ([]domain.PoProgressHeader, error) {
	datas, err := s.repository.FindAllProg(id)
	return datas, err
}
