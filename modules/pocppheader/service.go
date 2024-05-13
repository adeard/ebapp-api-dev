package pocppheader

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	FindCpp(id string) (domain.PoCppHeader, error)
	FindCppByRunNumProgress(id string) (domain.PoCppHeader, error)
	FindAllProg(id string) ([]domain.PoCppHeader, error)
	Delete(id string) error
	Update(id string, input domain.PoCppHeaderUpdate) (domain.PoCppHeader, error)
	UpdateStatus(id string, status string) (domain.PoCppHeader, error)
	Store(input domain.PoCppHeader) (domain.PoCppHeader, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) FindCpp(id string) (domain.PoCppHeader, error) {
	data, err := s.repository.FindCpp(id)
	return data, err
}

func (s *service) FindCppByRunNumProgress(id string) (domain.PoCppHeader, error) {
	data, err := s.repository.FindCppByRunNumProgress(id)
	return data, err
}

func (s *service) FindAllProg(id string) ([]domain.PoCppHeader, error) {
	datas, err := s.repository.FindAllProg(id)
	return datas, err
}

func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)
	return err
}

func (s *service) Update(id string, input domain.PoCppHeaderUpdate) (domain.PoCppHeader, error) {
	data, err := s.repository.Update(id, input)
	return data, err
}

func (s *service) UpdateStatus(id string, status string) (domain.PoCppHeader, error) {
	data, err := s.repository.UpdateStatus(id, status)
	return data, err
}

func (s *service) Store(input domain.PoCppHeader) (domain.PoCppHeader, error) {
	data, err := s.repository.Store(input)
	if err == nil {
		s.repository.CloneProgress(input.RunNumProgress, input.RunNum)
	}
	return data, err
}
