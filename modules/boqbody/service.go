package boqbody

import (
	"ebapp-api-dev/domain"
	"strconv"
)

type Service interface {
	GetAll(input domain.BoqBodyRequest) ([]domain.BoqBody, error)
	GetByRunNum(runNum string) ([]domain.BoqBody, error)
	GetByParentId(parentId string) ([]domain.BoqBody, error)
	FindByItemNo(itemNo string) (domain.BoqBody, error)
	Store(input domain.BoqBody) (domain.BoqBody, error)
	Update(input domain.BoqBody, id string) (domain.BoqBody, error)
	DeleteByID(id int, runNum string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.BoqBodyRequest) ([]domain.BoqBody, error) {
	boqBody, err := s.repository.FindAll(input)
	return boqBody, err
}

func (s *service) GetByRunNum(runNum string) ([]domain.BoqBody, error) {
	boqBody, err := s.repository.FindByRunNum(runNum)
	return boqBody, err
}

func (s *service) GetByParentId(parentId string) ([]domain.BoqBody, error) {
	boqBody, err := s.repository.FindByParentID(parentId)
	return boqBody, err
}

func (s *service) Store(input domain.BoqBody) (domain.BoqBody, error) {
	boqBody, err := s.repository.Store(input)
	return boqBody, err
}

func (s *service) Update(input domain.BoqBody, id string) (domain.BoqBody, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return domain.BoqBody{}, err
	}

	boqBodies, err := s.repository.FindByRunNum(input.RunNum)
	if err != nil {
		return boqBodies[0], err
	}

	boqBody := boqBodies[0]
	finalUpdateBoqBody := domain.BoqBody{
		Id:                intId,
		ParentId:          input.ParentId,
		RunNum:            boqBody.RunNum,
		ItemNo:            input.ItemNo,
		ItemLevel:         boqBody.ItemLevel,
		ItemDescription:   input.ItemDescription,
		ItemSpecification: input.ItemSpecification,
		Qty:               input.Qty,
		Unit:              input.Unit,
		Price:             input.Price,
		Currency:          input.Currency,
		Note:              input.Note,
	}

	result, err := s.repository.Update(finalUpdateBoqBody)
	return result, err
}

func (s *service) FindByItemNo(itemNo string) (domain.BoqBody, error) {
	boqBody, err := s.repository.FindByItemNo(itemNo)
	return boqBody, err
}

func (s *service) DeleteByID(id int, runNum string) error {
	// Cek terlebih dahulu apakah data dengan ID tersebut ada atau tidak
	_, err := s.repository.FindById(strconv.Itoa(id))
	if err != nil {
		// Jika data tidak ditemukan, kembalikan error
		return err
	}

	// Panggil fungsi DeleteByID dari repository untuk menghapus data BoQ Body berdasarkan ID
	err = s.repository.DeleteByID(strconv.Itoa(id), runNum)
	if err != nil {
		// Jika ada kesalahan saat menghapus, tangani sesuai kebutuhan (misalnya kembalikan pesan kesalahan)
		return err
	}

	return nil
}
