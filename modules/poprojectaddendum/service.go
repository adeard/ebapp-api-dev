package poprojectaddendum

import "ebapp-api-dev/domain"

type Service interface {
	GetByPo(po string) ([]domain.PoProjectAddendum, error)
	Delete(id string, po string, item string) error
	Store(input domain.PoProjectAddendum) (domain.PoProjectAddendum, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetByPo(po string) ([]domain.PoProjectAddendum, error) {
	poProjectAddendum, err := s.repository.FindByPo(po)
	return poProjectAddendum, err
}

func (s *service) Delete(id string, po string, item string) error {
	err := s.repository.Delete(id, po, item)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Store(input domain.PoProjectAddendum) (domain.PoProjectAddendum, error) {
	poAddendum, err := s.repository.Store(input)
	return poAddendum, err
}
