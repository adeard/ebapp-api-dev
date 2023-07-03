package boqbody

import "ebapp-api-dev/domain"

type Service interface {
	GetAll(input domain.BoqBodyRequest) ([]domain.BoqBody, error)
	Store(input domain.BoqBodyRequest) (domain.BoqBodyRequest, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.BoqBodyRequest) ([]domain.BoqBody, error) {

	boqBody, err := s.repository.FindAll(input)

	return boqBody, err
}

func (s *service) Store(input domain.BoqBodyRequest) (domain.BoqBodyRequest, error) {
	boqBody, err := s.repository.Store(input)

	return boqBody, err
}
