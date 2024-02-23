package poprogressheaderaddendum

import "ebapp-api-dev/domain"

type Service interface {
	FindAllProg(id string) ([]domain.PoProgressHeaderAddendum, error)
	Store(input domain.PoProgressHeaderAddendum) (domain.PoProgressHeaderAddendum, error)
	Update(id string, po string, input domain.PoProgressHeaderAddendumUpdate) (domain.PoProgressHeaderAddendumUpdate, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) FindAllProg(id string) ([]domain.PoProgressHeaderAddendum, error) {
	data, err := s.repository.FindAllProg(id)
	return data, err
}

func (s *service) Store(input domain.PoProgressHeaderAddendum) (domain.PoProgressHeaderAddendum, error) {
	data, err := s.repository.Store(input)
	return data, err
}

func (s *service) Update(id string, po string, input domain.PoProgressHeaderAddendumUpdate) (domain.PoProgressHeaderAddendumUpdate, error) {
	data, err := s.repository.Update(id, po, input)
	return data, err
}

func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)
	return err
}
