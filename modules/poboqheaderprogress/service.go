package poboqheaderprogress

import "ebapp-api-dev/domain"

type Service interface {
	Store(input domain.PoBoqHeaderProgress) (domain.PoBoqHeaderProgress, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
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
