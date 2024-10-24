package poboqbodycpp

import (
	"ebapp-api-dev/domain"
	"fmt"
	"strings"
)

type Service interface {
	GetByRunNum(runNum string, runNumProgress string, order string) ([]domain.PoBoqBodyCppProgress, error)
	CountByRunNum(runNum string) (int, error)
	SelectMaxOrder(runNum string) (int, error)
	Store(input domain.PoBoqBodyCpp) (domain.PoBoqBodyCpp, error)
	Update(runNum string, order string, mainId int, parentId int, status bool, note string) (domain.PoBoqBodyCppProgress, error)
	FindByItemNo(itemNo string) (domain.PoBoqBodyCpp, error)
	Delete(id string) error
	CloneCpp(previousRunNum string, nextRunNum string) error
	GroupItemsByParentServerSide(items []domain.PoBoqBodyCppResponse, parentId int) []domain.PoBoqBodyCppResponse
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

func (s *service) Update(runNum string, order string, mainId int, parentId int, status bool, note string) (domain.PoBoqBodyCppProgress, error) {
	updatedCpp, err := s.repository.Update(runNum, order, mainId, parentId, status, note)
	if err != nil {
		// Mengembalikan error jika terjadi kesalahan saat melakukan pembaruan
		return domain.PoBoqBodyCppProgress{}, err
	}

	return updatedCpp, nil
}

func (s *service) FindByItemNo(itemNo string) (domain.PoBoqBodyCpp, error) {
	boqBody, err := s.repository.FindByItemNo(itemNo)
	return boqBody, err
}

func (s *service) GetByRunNum(runNum string, runNumProgress string, order string) ([]domain.PoBoqBodyCppProgress, error) {
	poboqbody, err := s.repository.FindByRunNum(runNum, runNumProgress, order)
	return poboqbody, err
}

func (s *service) CountByRunNum(runNum string) (int, error) {
	total, err := s.repository.CountRunNum(runNum)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (s *service) SelectMaxOrder(runNum string) (int, error) {
	total, err := s.repository.SelectMaxOrder(runNum)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CloneCpp(previousRunNum string, nextRunNum string) error {
	//header
	var getExistCppHeader, errHeader = s.repository.GetByRunNumHeader(previousRunNum)
	if errHeader != nil {
		return errHeader
	}

	if len(getExistCppHeader) < 1 {
		parts := strings.Split(previousRunNum, "/")
		newStr := strings.Join(parts[:len(parts)-2], "/")
		getExistCppHeader, errHeader = s.repository.GetByRunNumBoqHeader(newStr)
	}

	newCppHeader := []domain.PoBoqHeaderCpp{}

	for _, bodyProgressDataHeader := range getExistCppHeader {
		tempCppHeader := bodyProgressDataHeader
		tempCppHeader.PekerjaanNo = nextRunNum
		newCppHeader = append(newCppHeader, tempCppHeader)
	}

	errHeader = s.repository.InsertBatchHeader(newCppHeader)
	if errHeader != nil {
		return errHeader
	}

	//body
	var getExistCpp, err = s.repository.GetByRunNum(previousRunNum)
	if err != nil {
		return err
	}

	if len(getExistCpp) < 1 {
		parts := strings.Split(previousRunNum, "/")
		newStr := strings.Join(parts[:len(parts)-2], "/")
		getExistCpp, err = s.repository.GetByRunNumBoqBody(newStr)
	}

	newCpp := []domain.PoBoqBodyCpp{}

	for _, bodyProgressData := range getExistCpp {
		tempCpp := bodyProgressData
		tempCpp.RunNum = nextRunNum
		tempCpp.Status = bodyProgressData.Status
		tempCpp.Note = bodyProgressData.Note

		newCpp = append(newCpp, tempCpp)
	}

	err = s.repository.InsertBatch(newCpp)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GroupItemsByParentServerSide(items []domain.PoBoqBodyCppResponse, parentId int) []domain.PoBoqBodyCppResponse {
	var result []domain.PoBoqBodyCppResponse

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
