package podatasapheader

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	GetTitle(po string) ([]domain.PoDataSapHeaderTitle, error)
	GetWbs(po string) ([]domain.ReadWbs, error)
	GetArea(id string) (domain.DataMasterPlant, error)
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

func (s *service) GetWbs(id string) ([]domain.ReadWbs, error) {
	wbs, err := s.repository.CheckWbs(id)
	return wbs, err
}

func (s *service) GetArea(id string) (domain.DataMasterPlant, error) {
	dataMasterPlant, err := s.repository.CheckArea(id)
	return dataMasterPlant, err
}
