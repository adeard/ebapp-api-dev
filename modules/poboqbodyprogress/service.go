package poboqbodyprogress

import "ebapp-api-dev/domain"

type Service interface {
	GetByRunNum(runNum string, order string) ([]domain.PoBoqBodyProgress, error)
	Store(input domain.PoBoqBodyProgress) (domain.PoBoqBodyProgress, error)
	FindByItemNo(itemNo string) (domain.PoBoqBodyProgress, error)
	Delete(id string) error
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

func (s *service) FindByItemNo(itemNo string) (domain.PoBoqBodyProgress, error) {
	boqBody, err := s.repository.FindByItemNo(itemNo)
	return boqBody, err
}

func (s *service) GetByRunNum(runNum string, order string) ([]domain.PoBoqBodyProgress, error) {
	poboqbody, err := s.repository.FindByRunNum(runNum, order)
	return poboqbody, err
}

func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
