package poboqbody

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/modules/listproject"
	"fmt"
	"math"
	"strconv"
)

type Service interface {
	GetByRunNum(runNum string, order string) ([]domain.PoBoqBody, error)
	CheckBoqBody(id string, order string, mainId string) ([]domain.PoBoqBody, error)
	Store(input domain.PoBoqBody) (domain.PoBoqBody, error)
	FindByItemNo(itemNo string) (domain.PoBoqBody, error)
	Delete(id string, order string, mainId string) error
	DeleteByOrder(id string, order string) error
	Update(input domain.PoBoqBody) (domain.PoBoqBody, error)
	CalculateByRunNumAndOrder(runNum string, order string) (float64, error)
	GroupItemsByParent(items []domain.PoBoqBodyResponse, parentId int) []domain.PoBoqBodyResponse
	FindLastId(runNum string, order string) (int, error)
	Adopth(oldRunNum string, newRunNum string, newOrder string) error
	GroupItemsByParentServerSide(items []domain.PoBoqBodyServerSideResponse, parentId int) []domain.PoBoqBodyServerSideResponse
}

type service struct {
	repository         Repository
	listProjectService listproject.Service
}

func NewService(repository Repository, listProjectService listproject.Service) *service {
	return &service{repository, listProjectService}
}

func (s *service) GetByRunNum(runNum string, order string) ([]domain.PoBoqBody, error) {
	poboqbody, err := s.repository.FindByRunNum(runNum, order)
	return poboqbody, err
}

func (s *service) Store(input domain.PoBoqBody) (domain.PoBoqBody, error) {
	poBoqBody, err := s.repository.Store(input)
	if err != nil {
		return poBoqBody, err
	}

	s.listProjectService.SyncCanProgressFalseService(poBoqBody.RunNum)
	return poBoqBody, err
}

func (s *service) Adopth(oldRunNum string, newRunNum string, newOrder string) error {
	err := s.repository.CopyBoqBodyToPoBoqBody(oldRunNum, newRunNum, newOrder)
	if err != nil {
		return err
	}
	s.listProjectService.SyncCanProgressFalseService(newRunNum)
	return err
}

func (s *service) FindByItemNo(itemNo string) (domain.PoBoqBody, error) {
	boqBody, err := s.repository.FindByItemNo(itemNo)
	return boqBody, err
}

func (s *service) FindLastId(runNum string, order string) (int, error) {
	value, err := s.repository.GenerateMainId(runNum, order)
	return value, err
}

func (s *service) Delete(id string, order string, mainId string) error {
	_, err := s.repository.FindBoq(id, order, mainId)
	if err != nil {
		return err
	}

	err = s.repository.Delete(id, order, mainId)
	if err != nil {
		return err
	}

	s.listProjectService.SyncCanProgressFalseService(id)

	return nil
}

func (s *service) DeleteByOrder(id string, order string) error {
	err := s.repository.DeleteByOrder(id, order)
	if err != nil {
		return err
	}

	s.listProjectService.SyncCanProgressFalseService(id)

	return nil
}

func (s *service) Update(input domain.PoBoqBody) (domain.PoBoqBody, error) {
	// Cek apakah item_no sudah ada
	isUnique, err := s.repository.CheckItemNo(input.RunNum, input.Order, input.ItemNo)
	if err != nil {
		return domain.PoBoqBody{}, err
	}

	// Jika item_no tidak unik, periksa apakah main_id yang sama
	if !isUnique {
		existingMainId, err := s.repository.SelectMainId(input.RunNum, input.Order, input.ItemNo)
		if err != nil {
			return domain.PoBoqBody{}, err
		}

		if existingMainId != input.Id {
			return domain.PoBoqBody{}, fmt.Errorf("ItemNo '%s' sudah digunakan oleh item lain", input.ItemNo)
		}
	}

	// Ambil data berdasarkan RunNum, Order, dan MainId
	result, err := s.repository.FindBoq(input.RunNum, input.Order, strconv.Itoa(input.Id))
	if err != nil {
		return domain.PoBoqBody{}, err
	}

	if len(result) == 0 {
		return domain.PoBoqBody{}, nil
	}

	poBoqBody := result[0]

	dataForUpdate := domain.PoBoqBody{
		Id:                poBoqBody.Id,
		ParentId:          poBoqBody.ParentId,
		RunNum:            poBoqBody.RunNum,
		Order:             poBoqBody.Order,
		ItemNo:            input.ItemNo,
		ItemLevel:         poBoqBody.ItemLevel,
		ItemDescription:   input.ItemDescription,
		ItemSpecification: input.ItemSpecification,
		Qty:               input.Qty,
		Unit:              input.Unit,
		Price:             input.Price,
		Currency:          input.Currency,
		Note:              input.Note,
		StartDate:         input.StartDate,
		EndDate:           input.EndDate,
		StartDateActual:   input.StartDateActual,
		EndDateActual:     input.EndDateActual,
	}

	poBoqBodies, err := s.repository.Update(dataForUpdate)

	if err != nil {
		return poBoqBodies, err
	}
	s.listProjectService.SyncCanProgressFalseService(poBoqBodies.RunNum)

	return poBoqBodies, err
}

func (s *service) CheckBoqBody(id string, order string, mainId string) ([]domain.PoBoqBody, error) {
	data, err := s.repository.FindBoq(id, order, mainId)
	return data, err
}

func (s *service) CalculateByRunNumAndOrder(runNum string, order string) (float64, error) {

	total := float64(0)

	poBoqBodyDatas, err := s.repository.GetByRunNumAndOrder(runNum, order)
	if err != nil {
		return 0, err
	}

	if len(poBoqBodyDatas) == 0 {
		return total, nil
	}

	for _, poBoqBodyData := range poBoqBodyDatas {
		total += float64(poBoqBodyData.Qty) * float64(poBoqBodyData.Price)
	}

	return math.Round(total), err
}

func (s *service) GroupItemsByParent(items []domain.PoBoqBodyResponse, parentId int) []domain.PoBoqBodyResponse {
	var result []domain.PoBoqBodyResponse

	for _, item := range items {
		if item.ParentId == parentId {
			children := s.GroupItemsByParent(items, item.Id)
			item.Children = children
			result = append(result, item)
		}
	}

	return result
}

func (s *service) GroupItemsByParentServerSide(items []domain.PoBoqBodyServerSideResponse, parentId int) []domain.PoBoqBodyServerSideResponse {
	var result []domain.PoBoqBodyServerSideResponse

	for _, item := range items {
		if item.ParentId == parentId {
			concatenatedOrder := item.Order + "-" + fmt.Sprint(item.Id)
			item.Order = concatenatedOrder
			children := s.GroupItemsByParentServerSide(items, item.Id)
			item.Children = children
			result = append(result, item)
		}
	}

	return result
}
