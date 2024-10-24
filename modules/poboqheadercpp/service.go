package poboqheadercpp

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/modules/poboqbodycpp"
	"strings"
)

type Service interface {
	GetCpp(id string) ([]domain.PoBoqHeaderCpp, error)
	Store(input domain.PoBoqHeaderCpp) (domain.PoBoqHeaderCpp, error)
	Delete(id string) error
	GetBodyServerSide(id string, page int, pageSize int, item string, desc string) ([]domain.PoBoqHeaderCppWithBodyServerSide, error)
	GetBodyServerSideCount(id string, item string, desc string) (int64, error)
}

type service struct {
	repository          Repository
	poboqbodyCppService poboqbodycpp.Service
}

func NewService(repository Repository, poboqbodyCppService poboqbodycpp.Service) *service {
	return &service{repository, poboqbodyCppService}
}

func (s *service) GetCpp(id string) ([]domain.PoBoqHeaderCpp, error) {
	headers, err := s.repository.FindCpp(id)
	return headers, err
}

func (s *service) Store(input domain.PoBoqHeaderCpp) (domain.PoBoqHeaderCpp, error) {
	headers, err := s.repository.Store(input)
	return headers, err
}

func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetBodyServerSide(id string, page int, pageSize int, item string, desc string) ([]domain.PoBoqHeaderCppWithBodyServerSide, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	headers, err := s.repository.FindCppWithPagingServerSide(id, page, pageSize, item, desc)

	result := []domain.PoBoqHeaderCppWithBodyServerSide{}

	for _, headersData := range headers {
		parts := strings.Split(headersData.PekerjaanNo, "/")
		newStr := strings.Join(parts[:len(parts)-1], "/")
		poBoqBody, err := s.poboqbodyCppService.GetByRunNum(headersData.PekerjaanNo, newStr, headersData.Order)

		if err != nil {
			return nil, err
		}

		var poBoqBodyResponse []domain.PoBoqBodyCppResponse
		for _, body := range poBoqBody {
			poBoqBodyResponse = append(poBoqBodyResponse, domain.PoBoqBodyCppResponse{
				Id:                body.Id,
				ParentId:          body.ParentId,
				RunNum:            body.RunNum,
				Order:             body.Order,
				ItemNo:            body.ItemNo,
				ItemLevel:         body.ItemLevel,
				ItemDescription:   body.ItemDescription,
				ItemSpecification: body.ItemSpecification,
				Qty:               body.Qty,
				Status:            body.Status,
				Note:              body.Note,
				Unit:              body.Unit,
				Price:             body.Price,
				Currency:          body.Currency,
				PreviousVolume:    body.PreviousVolume,
				CurrentVolume:     body.CurrentVolume,
			})
		}

		result = append(result, domain.PoBoqHeaderCppWithBodyServerSide{
			PoBoqHeaderCpp: headersData,
			BoqBodyCpp:     s.poboqbodyCppService.GroupItemsByParentServerSide(poBoqBodyResponse, 0),
		})
	}

	return result, err
}

func (s *service) GetBodyServerSideCount(id string, item string, desc string) (int64, error) {
	result, err := s.repository.FindCppWithPagingServerSideCount(id, item, desc)
	return result, err
}
