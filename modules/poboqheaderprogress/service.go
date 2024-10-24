package poboqheaderprogress

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/modules/poboqbodyprogress"
)

type Service interface {
	GetProgress(id string) ([]domain.PoBoqHeaderProgress, error)
	Store(input domain.PoBoqHeaderProgress) (domain.PoBoqHeaderProgress, error)
	Delete(id string) error
	GetBodyServerSide(id string, page int, pageSize int, item string, desc string) ([]domain.PoBoqHeaderProgressWithBodyServerSide, error)
	GetBodyServerSideCount(id string, item string, desc string) (int64, error)
	SyncProgress(run_num string) error
}

type service struct {
	repository               Repository
	poboqbodyProgressService poboqbodyprogress.Service
}

func NewService(repository Repository, poboqbodyProgressService poboqbodyprogress.Service) *service {
	return &service{repository, poboqbodyProgressService}
}

func (s *service) GetProgress(id string) ([]domain.PoBoqHeaderProgress, error) {
	headers, err := s.repository.FindProgress(id)
	return headers, err
}

func (s *service) Store(input domain.PoBoqHeaderProgress) (domain.PoBoqHeaderProgress, error) {
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

func (s *service) GetBodyServerSide(id string, page int, pageSize int, item string, desc string) ([]domain.PoBoqHeaderProgressWithBodyServerSide, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	headers, err := s.repository.FindProgressWithPagingServerSide(id, page, pageSize, item, desc)

	result := []domain.PoBoqHeaderProgressWithBodyServerSide{}

	for _, headersData := range headers {
		poBoqBody, err := s.poboqbodyProgressService.GetByRunNum(headersData.PekerjaanNo, headersData.Order)

		if err != nil {
			return nil, err
		}

		var poBoqBodyResponse []domain.PoBoqBodyProgressResponse
		for _, body := range poBoqBody {
			//concatenatedOrder := body.Order + "-" + fmt.Sprint(body.Id)
			poBoqBodyResponse = append(poBoqBodyResponse, domain.PoBoqBodyProgressResponse{
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
				PreviousVolume:    body.PreviousVolume,
				CurrentVolume:     body.CurrentVolume,
			})
		}

		result = append(result, domain.PoBoqHeaderProgressWithBodyServerSide{
			PoBoqHeaderProgress: headersData,
			BoqBodyProgress:     s.poboqbodyProgressService.GroupItemsByParentServerSide(poBoqBodyResponse, 0),
		})

	}

	return result, err
}

func (s *service) GetBodyServerSideCount(id string, item string, desc string) (int64, error) {
	result, err := s.repository.FindProgressWithPagingServerSideCount(id, item, desc)
	return result, err
}

func (s *service) SyncProgress(run_num string) error {
	err := s.repository.SyncProgress(run_num)
	return err
}
