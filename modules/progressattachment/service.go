package progressattachment

import "ebapp-api-dev/domain"

type Service interface {
	GetAttachment(id string) ([]domain.ProgressAttachment, error)
	Store(input domain.ProgressAttachment) (domain.ProgressAttachment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAttachment(id string) ([]domain.ProgressAttachment, error) {
	data, err := s.repository.GetAttachment(id)
	return data, err
}

func (s *service) Store(input domain.ProgressAttachment) (domain.ProgressAttachment, error) {
	data, err := s.repository.Store(input)
	return data, err
}
