package podatasapheader

import "ebapp-api-dev/domain"

type Service interface {
	GetByPo(po string) ([]domain.PoDataSapHeader, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetByPo(po string) ([]domain.PoDataSapHeader, error) {
	poDataSapHeader, err := s.repository.FindByPo(po)
	return poDataSapHeader, err
}
