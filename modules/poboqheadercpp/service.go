package poboqheadercpp

import "ebapp-api-dev/domain"

type Service interface {
	GetCpp(id string) ([]domain.PoBoqHeaderCpp, error)
	Store(input domain.PoBoqHeaderCpp) (domain.PoBoqHeaderCpp, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCpp(id string) ([]domain.PoBoqHeaderCpp, error) {
	headers, err := s.repository.FindCpp(id)
	return headers, err
}

func (s *service) Store(input domain.PoBoqHeaderCpp) (domain.PoBoqHeaderCpp, error) {
	headers, err := s.repository.Store(input)
	if err == nil {
		s.repository.CloneProgress(input.PekerjaanNoProgress, input.PekerjaanNo)
	}
	return headers, err
}

func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
