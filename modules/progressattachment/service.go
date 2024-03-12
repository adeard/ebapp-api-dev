package progressattachment

import "ebapp-api-dev/domain"

type Service interface {
	Store(input domain.ProgressAttachment) (domain.ProgressAttachment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Store(input domain.ProgressAttachment) (domain.ProgressAttachment, error) {
	data, err := s.repository.Store(input)
	return data, err
}
