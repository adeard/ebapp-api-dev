package poboqheader

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/modules/listproject"
	"ebapp-api-dev/modules/poboqbody"
	"math"
	"time"
)

type Service interface {
	GetByPekerjaanNo(id string) ([]domain.PoBoqHeader, error)
	Delete(id string, po string, order string) error
	Store(input domain.PoBoqHeader) (domain.PoBoqHeader, error)
	SyncActualPrice(pekerjaanNo string) error
	UpdateByPekerjaanNoAndRunNum(pekerjaanNo string, runNum string, input domain.PoBoqHeader) error
	GetByPekerjaanNoWithBody(id string, page int, pageSize int) ([]domain.PoBoqHeaderWithBody, error)
	GetByPekerjaanNoWithBodyByOrder(id string, order int) ([]domain.PoBoqHeaderWithBody, error)

	GetByPekerjaanNoWithBodyServerSide(id string, page int, pageSize int, item string, desc string) ([]domain.PoBoqHeaderWithBodyServerSide, error)
	GetByPekerjaanNoWithBodyServerSideCount(id string, item string, desc string) (int64, error)
}

type service struct {
	repository         Repository
	poboqbodyService   poboqbody.Service
	listProjectService listproject.Service
}

func NewService(repository Repository, poboqbodyService poboqbody.Service, listProjectService listproject.Service) *service {
	return &service{repository, poboqbodyService, listProjectService}
}

func (s *service) GetByPekerjaanNo(id string) ([]domain.PoBoqHeader, error) {
	headers, err := s.repository.FindByPekerjaanNo(id)
	return headers, err
}

func (s *service) Store(input domain.PoBoqHeader) (domain.PoBoqHeader, error) {
	headers, err := s.repository.Store(input)
	if err != nil {
		return headers, err
	}

	// Panggil SyncCanProgressFalseService dengan pekerjaanNo dari input
	err = s.listProjectService.SyncCanProgressFalseService(input.PekerjaanNo)
	if err != nil {
		return headers, err
	}

	return headers, err
}

func (s *service) Delete(id string, po string, order string) error {
	err := s.repository.Delete(id, po, order)
	if err != nil {
		return err
	}

	// Panggil SyncCanProgressFalseService dengan id (PekerjaanNo)
	err = s.listProjectService.SyncCanProgressFalseService(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SyncActualPrice(pekerjaanNo string) error {
	canProgress := 1
	headers, err := s.GetByPekerjaanNo(pekerjaanNo)
	if err != nil {
		return err
	}

	for _, headersData := range headers {
		totalPrice, err := s.poboqbodyService.CalculateByRunNumAndOrder(headersData.PekerjaanNo, headersData.Order)
		if err != nil {
			return err
		}

		err = s.UpdateByPekerjaanNoAndRunNum(headersData.PekerjaanNo, headersData.Order, domain.PoBoqHeader{ActualPrice: totalPrice})

		if err != nil {
			return err
		}

		if math.Round(float64(headersData.Qty)*float64(headersData.Price)) != totalPrice {
			canProgress = 0
		}
	}

	err = s.listProjectService.UpdateByPekerjaanNo(pekerjaanNo, domain.ListProject{CanProgress: canProgress})

	return err
}

func (s *service) UpdateByPekerjaanNoAndRunNum(pekerjaanNo string, runNum string, input domain.PoBoqHeader) error {
	updateData := map[string]interface{}{}

	//loc, _ := time.LoadLocation("Asia/Jakarta")

	// if input.ActualPrice > 0 {
	// 	updateData["actual_price"] = input.ActualPrice
	// }

	updateData["actual_price"] = input.ActualPrice

	updateData["last_updated"] = time.Now().UTC()

	err := s.repository.UpdateByPekerjaanNoAndRunNum(pekerjaanNo, runNum, updateData)

	return err
}

// Raffi -- Pagination BOQ Header with body

func (s *service) GetByPekerjaanNoWithBody(id string, page int, pageSize int) ([]domain.PoBoqHeaderWithBody, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	headers, err := s.repository.FindByPekerjaanNoWithPaging(id, page, pageSize)

	result := []domain.PoBoqHeaderWithBody{}

	for _, headersData := range headers {
		poBoqBody, err := s.poboqbodyService.GetByRunNum(headersData.PekerjaanNo, headersData.Order)

		if err != nil {
			return nil, err
		}

		var poBoqBodyResponse []domain.PoBoqBodyResponse
		for _, body := range poBoqBody {
			poBoqBodyResponse = append(poBoqBodyResponse, domain.PoBoqBodyResponse{
				Id:                body.Id,
				ParentId:          body.ParentId,
				RunNum:            body.RunNum,
				Order:             body.Order,
				ItemNo:            body.ItemNo,
				ItemLevel:         body.ItemLevel,
				ItemDescription:   body.ItemDescription,
				ItemSpecification: body.ItemSpecification,
				Qty:               body.Qty,
				Unit:              body.Unit,
				Price:             body.Price,
				Currency:          body.Currency,
				Note:              body.Note,
				StartDate:         body.StartDate,
				EndDate:           body.EndDate,
				StartDateActual:   body.StartDateActual,
				EndDateActual:     body.EndDateActual,
			})
		}

		result = append(result, domain.PoBoqHeaderWithBody{
			PoBoqHeader: headersData,
			BoqBody:     s.poboqbodyService.GroupItemsByParent(poBoqBodyResponse, 0),
		})

	}

	return result, err
}

func (s *service) GetByPekerjaanNoWithBodyServerSide(id string, page int, pageSize int, item string, desc string) ([]domain.PoBoqHeaderWithBodyServerSide, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	headers, err := s.repository.FindByPekerjaanNoWithPagingServerSide(id, page, pageSize, item, desc)

	result := []domain.PoBoqHeaderWithBodyServerSide{}

	for _, headersData := range headers {
		poBoqBody, err := s.poboqbodyService.GetByRunNum(headersData.PekerjaanNo, headersData.Order)

		if err != nil {
			return nil, err
		}

		var poBoqBodyResponse []domain.PoBoqBodyServerSideResponse
		for _, body := range poBoqBody {
			//concatenatedOrder := body.Order + "-" + fmt.Sprint(body.Id)
			poBoqBodyResponse = append(poBoqBodyResponse, domain.PoBoqBodyServerSideResponse{
				Id:                body.Id,
				ParentId:          body.ParentId,
				RunNum:            body.RunNum,
				Order:             body.Order,
				ItemNo:            body.ItemNo,
				ItemLevel:         body.ItemLevel,
				ItemDescription:   body.ItemDescription,
				ItemSpecification: body.ItemSpecification,
				Qty:               body.Qty,
				Unit:              body.Unit,
				Price:             body.Price,
				Currency:          body.Currency,
				Note:              body.Note,
				StartDate:         body.StartDate,
				EndDate:           body.EndDate,
				StartDateActual:   body.StartDateActual,
				EndDateActual:     body.EndDateActual,
			})
		}

		result = append(result, domain.PoBoqHeaderWithBodyServerSide{
			PoBoqHeader: headersData,
			BoqBody:     s.poboqbodyService.GroupItemsByParentServerSide(poBoqBodyResponse, 0),
		})

	}

	return result, err
}

func (s *service) GetByPekerjaanNoWithBodyServerSideCount(id string, item string, desc string) (int64, error) {
	result, err := s.repository.FindByPekerjaanNoWithPagingServerSideCount(id, item, desc)
	return result, err
}

func (s *service) GetByPekerjaanNoWithBodyByOrder(id string, order int) ([]domain.PoBoqHeaderWithBody, error) {
	// Ambil data PoBoqHeader berdasarkan pekerjaan_no dan order
	headers, err := s.repository.FindByPekerjaanNoWithOrder(id, order)
	if err != nil {
		return nil, err
	}

	// Array untuk menyimpan hasil
	result := []domain.PoBoqHeaderWithBody{}

	// Iterasi setiap header
	for _, headersData := range headers {
		// Ambil data BoqBody berdasarkan pekerjaan_no dan order
		poBoqBody, err := s.poboqbodyService.GetByRunNum(headersData.PekerjaanNo, headersData.Order)
		if err != nil {
			return nil, err
		}

		// Konversi ke response BoqBody
		var poBoqBodyResponse []domain.PoBoqBodyResponse
		for _, body := range poBoqBody {
			poBoqBodyResponse = append(poBoqBodyResponse, domain.PoBoqBodyResponse{
				Id:                body.Id,
				ParentId:          body.ParentId,
				RunNum:            body.RunNum,
				Order:             body.Order,
				ItemNo:            body.ItemNo,
				ItemLevel:         body.ItemLevel,
				ItemDescription:   body.ItemDescription,
				ItemSpecification: body.ItemSpecification,
				Qty:               body.Qty,
				Unit:              body.Unit,
				Price:             body.Price,
				Currency:          body.Currency,
				Note:              body.Note,
				StartDate:         body.StartDate,
				EndDate:           body.EndDate,
				StartDateActual:   body.StartDateActual,
				EndDateActual:     body.EndDateActual,
			})
		}

		// Masukkan header dan BoqBody yang telah dikelompokkan ke dalam hasil
		result = append(result, domain.PoBoqHeaderWithBody{
			PoBoqHeader: headersData,
			BoqBody:     s.poboqbodyService.GroupItemsByParent(poBoqBodyResponse, 0),
		})
	}

	// Kembalikan hasil akhir
	return result, nil
}
