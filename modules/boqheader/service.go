package boqheader

import (
	"ebapp-api-dev/domain"
	"time"
)

type Service interface {
	GetAll(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error)
	GetActive(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error)
	GetByID(id string) (domain.BoqHeader, error)
	Store(input domain.BoqHeader) (domain.BoqHeader, error)
	Update(input domain.BoqHeader, id string) (domain.BoqHeader, error)
	Clone(oldId string, newId string) (string, error)
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

func (s *service) GetActive(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error) {
	boqHeaders, err := s.repository.FindAllActive(input)
	return boqHeaders, err
}

func (s *service) GetByID(id string) (domain.BoqHeader, error) {
	boqheader, err := s.repository.FindById(id)
	return boqheader, err
}

func (s *service) Store(input domain.BoqHeader) (domain.BoqHeader, error) {
	boqHeader, err := s.repository.Store(input)
	return boqHeader, err
}

func (s *service) Update(input domain.BoqHeader, id string) (domain.BoqHeader, error) {
	boqheader, err := s.repository.FindById(id)
	if err != nil {
		return boqheader, err
	}

	finalUpdateBoqHeader := domain.BoqHeader{
		RunNum:            boqheader.RunNum,
		BoqNo:             input.BoqNo,
		HeaderDescription: input.HeaderDescription,
		HeaderVersion:     input.HeaderVersion,
		HeaderStatus:      input.HeaderStatus,
		Created:           boqheader.Created,
		CreatedBy:         boqheader.CreatedBy,
		LastUpdated:       time.Now(),
		LastUpdatedBy:     input.LastUpdatedBy,
		Category:          input.Category,
		Remarks:           input.Remarks,
	}

	boqHeaders, err := s.repository.Update(finalUpdateBoqHeader)
	return boqHeaders, err
}

func (s *service) Clone(oldId string, newId string) (string, error) {
	data, err := s.repository.Clone(oldId, newId)
	return data, err
}
