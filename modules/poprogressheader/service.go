package poprogressheader

type Service interface {
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}
