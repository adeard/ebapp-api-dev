package popic

import "ebapp-api-dev/domain"

type Service interface {
	FindPicByPo(po string) ([]domain.PoPic, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) FindPicByPo(po string) ([]domain.PoPic, error) {
	pic, err := s.repository.FindByPo(po)
	return pic, err
}
