package popic

import "ebapp-api-dev/domain"

type Service interface {
	FindPic(uid string, po string, role string) (domain.PoPic, error)
	FindPicByLevel(uid string, po string, level int) (domain.PoPic, error)
	FindPicByPo(po string) ([]domain.PoPic, error)
	Store(input domain.PoPic) (domain.PoPic, error)
	Delete(uid string, po string, level int) (domain.PoPic, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) FindPic(uid string, po string, role string) (domain.PoPic, error) {
	pic, err := s.repository.FindPic(uid, po, role)
	return pic, err
}

func (s *service) FindPicByLevel(uid string, po string, level int) (domain.PoPic, error) {
	pic, err := s.repository.FindPicByLevel(uid, po, level)
	return pic, err
}

func (s *service) FindPicByPo(po string) ([]domain.PoPic, error) {
	pic, err := s.repository.FindByPo(po)
	return pic, err
}

func (s *service) Store(input domain.PoPic) (domain.PoPic, error) {
	pic, err := s.repository.Store(input)
	return pic, err
}

func (s *service) Delete(uid string, po string, level int) (domain.PoPic, error) {
	pic, err := s.repository.Delete(uid, po, level)
	return pic, err
}
