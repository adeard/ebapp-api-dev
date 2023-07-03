package auth

import "ebapp-api-dev/domain"

type Service interface {
	AuthTest() (domain.Auth, error)
	SignIn(authinput domain.AuthRequest) (domain.AuthResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AuthTest() (domain.Auth, error) {

	auth, err := s.repository.AuthTest()

	return auth, err
}

func (s *service) SignIn(authinput domain.AuthRequest) (domain.AuthResponse, error) {

	auth, err := s.repository.SignIn(domain.Auth(authinput))

	return auth, err
}
