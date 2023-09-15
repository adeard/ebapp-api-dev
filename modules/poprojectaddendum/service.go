package poprojectaddendum

import "ebapp-api-dev/domain"

type Service interface {
	GetByPo(po string) ([]domain.PoProjectAddendum, error)
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
