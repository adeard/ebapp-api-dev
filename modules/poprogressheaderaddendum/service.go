package poprogressheaderaddendum

type Service interface {
	Delete(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)
	return err
}
