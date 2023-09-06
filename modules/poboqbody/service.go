package poboqbody

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	GetByRunNum(runNum string) ([]domain.PoBoqBody, error)
	Store(input domain.PoBoqBody) (domain.PoBoqBody, error)
	FindByItemNo(itemNo string) (domain.PoBoqBody, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetByRunNum(runNum string) ([]domain.PoBoqBody, error) {
	poboqbody, err := s.repository.FindByRunNum(runNum)
	return poboqbody, err
}

func (s *service) Store(input domain.PoBoqBody) (domain.PoBoqBody, error) {
	poBoqBody, err := s.repository.Store(input)
	return poBoqBody, err
}

func (s *service) FindByItemNo(itemNo string) (domain.PoBoqBody, error) {
	boqBody, err := s.repository.FindByItemNo(itemNo)
	return boqBody, err
}
