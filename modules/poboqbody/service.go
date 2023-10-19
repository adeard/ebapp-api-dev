package poboqbody

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	GetByRunNum(runNum string, order string) ([]domain.PoBoqBody, error)
	Store(input domain.PoBoqBody) (domain.PoBoqBody, error)
	FindByItemNo(itemNo string) (domain.PoBoqBody, error)
	Delete(id string, order string, mainId string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetByRunNum(runNum string, order string) ([]domain.PoBoqBody, error) {
	poboqbody, err := s.repository.FindByRunNum(runNum, order)
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

func (s *service) Delete(id string, order string, mainId string) error {
	// Cek terlebih dahulu apakah data dengan ID tersebut ada atau tidak
	_, err := s.repository.FindBoq(id, order, mainId)
	if err != nil {
		return err
	}

	err = s.repository.Delete(id, order, mainId)
	if err != nil {
		return err
	}

	return nil
}
