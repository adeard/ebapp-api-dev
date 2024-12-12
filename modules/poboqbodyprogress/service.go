package poboqbodyprogress

import (
	"ebapp-api-dev/domain"
	"fmt"
	"strings"
)

type Service interface {
	GetByRunNum(runNum string, order string) ([]domain.PoBoqBodyProgress, error)
	CountByRunNum(runNum string) (int, error)
	SelectMaxOrder(runNum string) (int, error)
	Store(input domain.PoBoqBodyProgress) (domain.PoBoqBodyProgress, error)
	Update(runNum string, order string, mainId int, parentId int, current_volume float64) (domain.PoBoqBodyProgress, error)
	FindByItemNo(itemNo string) (domain.PoBoqBodyProgress, error)
	Delete(id string) error
	CloneProgress(previousProgress string, nextProgress string) error
	GroupItemsByParentServerSide(items []domain.PoBoqBodyProgressResponse, parentId int) []domain.PoBoqBodyProgressResponse
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input domain.PoBoqBodyProgress) (domain.PoBoqBodyProgress, error) {
	poBoqBodyProgress, err := s.repository.Store(input)
	return poBoqBodyProgress, err
}

func (s *service) Update(runNum string, order string, mainId int, parentId int, current_volume float64) (domain.PoBoqBodyProgress, error) {
	updatedProgress, err := s.repository.Update(runNum, order, mainId, parentId, current_volume)
	if err != nil {
		// Mengembalikan error jika terjadi kesalahan saat melakukan pembaruan
		return domain.PoBoqBodyProgress{}, err
	}

	return updatedProgress, nil
}

func (s *service) FindByItemNo(itemNo string) (domain.PoBoqBodyProgress, error) {
	boqBody, err := s.repository.FindByItemNo(itemNo)
	return boqBody, err
}

func (s *service) GetByRunNum(runNum string, order string) ([]domain.PoBoqBodyProgress, error) {
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

func (s *service) CloneProgress(previousProgress string, nextProgress string) error {
	var getExistProgress, err = s.repository.GetByRunNum(previousProgress)
	if err != nil {
		return err
	}

	if len(getExistProgress) < 1 {
		parts := strings.Split(previousProgress, "/")
		newStr := strings.Join(parts[:len(parts)-1], "/")
		getExistProgress, err = s.repository.GetByRunNumBoqBody(newStr)
		//return errors.New("Run num " + previousProgress + " not found")
	}

	//get is_addendum
	var getProgressAddendum, err1 = s.repository.GetByRunNumBoqHeaderProgressAddendum(previousProgress)
	if err1 != nil {
		return err
	}

	if len(getProgressAddendum) > 0 {
		parts1 := strings.Split(previousProgress, "/")
		newStr1 := strings.Join(parts1[:len(parts1)-1], "/")
		for _, gpa := range getProgressAddendum {
			getBoqBodyProgressAddendum, err3 := s.repository.GetByRunNumBoqBodyProgressAddendum(previousProgress, gpa.Order)
			if err3 == nil {
				if len(getBoqBodyProgressAddendum) == 0 {
					getAddendumBoqBody, err2 := s.repository.GetByRunNumBoqBodyAddendum(newStr1, gpa.Order)
					if err2 == nil {
						getExistProgress = append(getExistProgress, getAddendumBoqBody...)
					}
				}
			}
		}
	}

	newProgress := []domain.PoBoqBodyProgress{}

	for _, bodyProgressData := range getExistProgress {
		tempProgress := bodyProgressData
		tempProgress.RunNum = nextProgress
		tempProgress.PreviousVolume = bodyProgressData.CurrentVolume + bodyProgressData.PreviousVolume

		if bodyProgressData.CurrentVolume != 0 {
			tempProgress.CurrentVolume = 0
		}

		newProgress = append(newProgress, tempProgress)
	}

	err = s.repository.InsertBatch(newProgress)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GroupItemsByParentServerSide(items []domain.PoBoqBodyProgressResponse, parentId int) []domain.PoBoqBodyProgressResponse {
	var result []domain.PoBoqBodyProgressResponse

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
