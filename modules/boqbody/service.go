package boqbody

import (
	"ebapp-api-dev/domain"
	"fmt"
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

	GetByRunNumServerSide(runNum string, page int, pageSize int, item_no string, item_desc string) ([]domain.BoqBody, error)
	GetByRunNumServerSideCount(runNum string, item_no string, item_desc string) (int64, error)
	GetByParentIdServerSide(parentId string, run_num string) ([]domain.BoqBody, error)

	groupItemsByParentServerSide(items []domain.BoqBodyServerSide, parentId int) []domain.BoqBodyServerSide
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

func (s *service) GetByRunNumServerSide(runNum string, page int, pageSize int, item_no string, item_desc string) ([]domain.BoqBody, error) {
	boqBody, err := s.repository.FindByRunNumServerSide(runNum, page, pageSize, item_no, item_desc)
	return boqBody, err
}

func (s *service) GetByRunNumServerSideCount(runNum string, item_no string, item_desc string) (int64, error) {
	result, err := s.repository.FindByRunNumServerSideCount(runNum, item_no, item_desc)
	return result, err
}

func (s *service) GetByParentIdServerSide(parentId string, run_num string) ([]domain.BoqBody, error) {
	boqBody, err := s.repository.FindByParentIDServerSide(parentId, run_num)
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

func (s *service) groupItemsByParentServerSide(items []domain.BoqBodyServerSide, parentId int) []domain.BoqBodyServerSide {
	var result []domain.BoqBodyServerSide

	for _, item := range items {
		// if item.ParentId == parentId {
		_boqBody, _ := s.repository.FindByParentIDServerSide(strconv.Itoa(item.Id), item.RunNum)
		//item.Children = boqBody
		var _boqBodyServerSide []domain.BoqBodyServerSide
		for _, body := range _boqBody {
			_boqBodyServerSide = append(_boqBodyServerSide, domain.BoqBodyServerSide{
				Id:                body.Id,
				ParentId:          body.ParentId,
				RunNum:            body.RunNum,
				ItemNo:            body.ItemNo,
				ItemLevel:         body.ItemLevel,
				ItemDescription:   body.ItemDescription,
				ItemSpecification: body.ItemSpecification,
				Qty:               body.Qty,
				Unit:              body.Unit,
				Price:             body.Price,
				Currency:          body.Currency,
				Note:              body.Note,
				Order:             fmt.Sprint(item.Id) + "-" + strconv.Itoa(body.Id),
			})
		}
		children := s.groupItemsByParentServerSide(_boqBodyServerSide, 0)
		//var _boqBody []domain.BoqBody
		// for _, body := range children {
		// 	_boqBody = append(_boqBody, domain.BoqBody{
		// 		Id:                body.Id,
		// 		ParentId:          body.ParentId,
		// 		RunNum:            body.RunNum,
		// 		ItemNo:            body.ItemNo,
		// 		ItemLevel:         body.ItemLevel,
		// 		ItemDescription:   body.ItemDescription,
		// 		ItemSpecification: body.ItemSpecification,
		// 		Qty:               body.Qty,
		// 		Unit:              body.Unit,
		// 		Price:             body.Price,
		// 		Currency:          body.Currency,
		// 		Note:              body.Note,
		// 		Order:             strconv.Itoa(body.Id) + "-" + fmt.Sprint(item.Id),
		// 	})
		// }
		item.Children = children
		result = append(result, item)
		//}
	}

	return result
}
