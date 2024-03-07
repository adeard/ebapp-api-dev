package poprogressheader

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	FindProg(id string) (domain.PoProgressHeader, error)
	FindAllProg(id string) ([]domain.PoProgressHeader, error)
	Delete(id string) error
	Update(id string, input domain.PoProgressHeaderUpdate) (domain.PoProgressHeader, error)
	EbappUpdate(id string, input domain.PoProgressHeaderUpdateEbapp) (domain.PoProgressHeader, error)
	Store(input domain.PoProgressHeader) (domain.PoProgressHeader, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) FindProg(id string) (domain.PoProgressHeader, error) {
	data, err := s.repository.FindProg(id)
	return data, err
}

func (s *service) FindAllProg(id string) ([]domain.PoProgressHeader, error) {
	datas, err := s.repository.FindAllProg(id)
	return datas, err
}

func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)
	return err
}

func (s *service) Update(id string, input domain.PoProgressHeaderUpdate) (domain.PoProgressHeader, error) {
	data, err := s.repository.Update(id, input)
	return data, err
}

func (s *service) EbappUpdate(id string, input domain.PoProgressHeaderUpdateEbapp) (domain.PoProgressHeader, error) {
	data, err := s.repository.EbappUpdate(id, input)
	return data, err
}

func (s *service) Store(input domain.PoProgressHeader) (domain.PoProgressHeader, error) {
	data, err := s.repository.Store(input)
	return data, err
}
