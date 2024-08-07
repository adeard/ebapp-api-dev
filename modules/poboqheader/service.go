package poboqheader

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/modules/poboqbody"
	"fmt"
)

type Service interface {
	GetByPekerjaanNo(id string) ([]domain.PoBoqHeader, error)
	Delete(id string, po string, order string) error
	Store(input domain.PoBoqHeader) (domain.PoBoqHeader, error)
	SyncActualPrice(pekerjaanNo string) error
	UpdateByPekerjaanNoAndRunNum(pekerjaanNo string, runNum string, input domain.PoBoqHeader) error
}

type service struct {
	repository       Repository
	poboqbodyService poboqbody.Service
}

func NewService(repository Repository, poboqbodyService poboqbody.Service) *service {
	return &service{repository, poboqbodyService}
}

func (s *service) GetByPekerjaanNo(id string) ([]domain.PoBoqHeader, error) {
	headers, err := s.repository.FindByPekerjaanNo(id)
	return headers, err
}

func (s *service) Store(input domain.PoBoqHeader) (domain.PoBoqHeader, error) {
	headers, err := s.repository.Store(input)
	return headers, err
}

func (s *service) Delete(id string, po string, order string) error {
	err := s.repository.Delete(id, po, order)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SyncActualPrice(pekerjaanNo string) error {
	headers, err := s.GetByPekerjaanNo(pekerjaanNo)

	for _, headersData := range headers {
		totalPrice, err := s.poboqbodyService.CalculateByRunNumAndOrder(headersData.PekerjaanNo, headersData.Order)
		if err != nil {
			return err
		}

		s.UpdateByPekerjaanNoAndRunNum(headersData.PekerjaanNo, headersData.Order, domain.PoBoqHeader{ActualPrice: fmt.Sprintf("%d", totalPrice)})
	}

	return err
}

func (s *service) UpdateByPekerjaanNoAndRunNum(pekerjaanNo string, runNum string, input domain.PoBoqHeader) error {
	updateData := map[string]interface{}{}

	if input.ActualPrice != "" {
		updateData["actual_price"] = input.ActualPrice
	}

	err := s.repository.UpdateByPekerjaanNoAndRunNum(pekerjaanNo, runNum, updateData)

	return err
}
