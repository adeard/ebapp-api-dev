package poboqbodycpp

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	GetByRunNum(runNum string, order string) ([]domain.PoBoqBodyCpp, error)
	CountByRunNum(runNum string) (int, error)
	// SelectMaxOrder(runNum string) (int, error)
	Store(input domain.PoBoqBodyCpp) (domain.PoBoqBodyCpp, error)
	Update(runNum string, order string, mainId int, parentId int, status bool, note string) (domain.PoBoqBodyCpp, error)
	FindByItemNo(itemNo string) (domain.PoBoqBodyCpp, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input domain.PoBoqBodyCpp) (domain.PoBoqBodyCpp, error) {
	poBoqBodyCpp, err := s.repository.Store(input)
	return poBoqBodyCpp, err
}

func (s *service) Update(runNum string, order string, mainId int, parentId int, status bool, note string) (domain.PoBoqBodyCpp, error) {
	updatedCpp, err := s.repository.Update(runNum, order, mainId, parentId, status, note)
	if err != nil {
		// Mengembalikan error jika terjadi kesalahan saat melakukan pembaruan
		return domain.PoBoqBodyCpp{}, err
	}

	return updatedCpp, nil
}

func (s *service) FindByItemNo(itemNo string) (domain.PoBoqBodyCpp, error) {
	boqBody, err := s.repository.FindByItemNo(itemNo)
	return boqBody, err
}

func (s *service) GetByRunNum(runNum string, order string) ([]domain.PoBoqBodyCpp, error) {
	poboqbody, err := s.repository.FindByRunNum(runNum, order)
	return poboqbody, err
}

func (s *service) CountByRunNum(runNum string) (int, error) {
	total, err := s.repository.CountRunNum(runNum)
	if err != nil {
		return 0, err
	}

	return total, nil
}

// func (s *service) SelectMaxOrder(runNum string) (int, error) {
// 	total, err := s.repository.SelectMaxOrder(runNum)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return total, nil
// }

func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
