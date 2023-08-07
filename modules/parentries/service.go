package parentries

import "ebapp-api-dev/domain"

type Service interface {
	GetAll(input domain.ParEntriesRequest) ([]domain.ParEntries, error)
	GetByID(id string) ([]domain.ParEntries, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll(input domain.ParEntriesRequest) ([]domain.ParEntries, error) {
	parEntriess, err := s.repository.FindAll(input)
	return parEntriess, err
}

func (s *service) GetByID(id string) ([]domain.ParEntries, error) {
	parEntries, err := s.repository.FindAllById(id)
	return parEntries, err
}
