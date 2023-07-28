package boqbody

import "ebapp-api-dev/domain"

type Service interface {
	GetAll(input domain.BoqBodyRequest) ([]domain.BoqBody, error)
	GetByID(runNum string) ([]domain.BoqBody, error)
	Store(input domain.BoqBody) (domain.BoqBody, error)
	Update(input domain.BoqBody, id string) (domain.BoqBody, error)
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

func (s *service) GetByID(runNum string) ([]domain.BoqBody, error) {
	boqBody, err := s.repository.FindByID(runNum)
	return boqBody, err
}

func (s *service) Store(input domain.BoqBody) (domain.BoqBody, error) {
	boqBody, err := s.repository.Store(input)
	return boqBody, err
}

func (s *service) Update(input domain.BoqBody, id string) (domain.BoqBody, error) {
	boqBody, err := s.repository.FindById(id)
	if err != nil {
		return boqBody, err
	}

	finalUpdateBoqBody := domain.BoqBody{
		Id:                boqBody.Id,
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

	boqBodies, err := s.repository.Update(finalUpdateBoqBody)
	return boqBodies, err
}
