package podatasapheader

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	GetTitle(po string) ([]domain.PoDataSapHeaderTitle, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetTitle(id string) ([]domain.PoDataSapHeaderTitle, error) {
	poDataSapHeaderTitle, err := s.repository.CheckTitle(id)
	return poDataSapHeaderTitle, err
}
