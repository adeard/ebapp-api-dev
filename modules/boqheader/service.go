package boqheader

import "ebapp-api-dev/domain"

type Service interface {
	GetAll(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error)
	Store(input domain.BoqHeader) (domain.BoqHeader, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error) {
	boqHeaders, err := s.repository.FindAll(input)
	return boqHeaders, err
}

func (s *service) Store(input domain.BoqHeader) (domain.BoqHeader, error) {
	boqHeader, err := s.repository.Store(input)
	return boqHeader, err
}
