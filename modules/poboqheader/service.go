package poboqheader

import "ebapp-api-dev/domain"

type Service interface {
	GetByPekerjaanNo(id string) ([]domain.PoBoqHeader, error)
	Store(input domain.PoBoqHeader) (domain.PoBoqHeader, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetByPekerjaanNo(id string) ([]domain.PoBoqHeader, error) {
	headers, err := s.repository.FindByPekerjaanNo(id)
	return headers, err
}

func (s *service) Store(input domain.PoBoqHeader) (domain.PoBoqHeader, error) {
	headers, err := s.repository.Store(input)
	return headers, err
}
