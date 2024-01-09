package poboqbodyprogress

import "ebapp-api-dev/domain"

type Service interface {
	Store(input domain.PoBoqBodyProgress) (domain.PoBoqBodyProgress, error)
	FindByItemNo(itemNo string) (domain.PoBoqBodyProgress, error)
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
