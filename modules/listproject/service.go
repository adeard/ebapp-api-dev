package listproject

import (
	"ebapp-api-dev/domain"
)

type Service interface {
	GetAll(input domain.ListProjectRequest) ([]domain.ListProject, error)
	GetByID(id string) (domain.ListProject, error)
	GetByPlant(ids []string) ([]domain.ListProject, error)
	Store(input domain.ListProject) (domain.ListProject, error)
	Store2(input domain.ListProject2) (domain.ListProject2, error)
	Store3(input domain.ListProject3) (domain.ListProject3, error)
	Store4(input domain.ListProject4) (domain.ListProject4, error)

	UpdateStatus(input domain.UpdateStatus, Id string) (domain.UpdateStatus, error)
	UpdateSpkNRetensi(input domain.UpdateSpkNRetensi, Id string) (domain.UpdateSpkNRetensi, error)
	UpdatePlanningActualDate(input domain.ModelUpdateActualPlanningDate) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.ListProjectRequest) ([]domain.ListProject, error) {
	listProjects, err := s.repository.FindAll(input)
	return listProjects, err
}

func (s *service) GetByID(id string) (domain.ListProject, error) {
	project, err := s.repository.FindById(id)
	return project, err
}

func (s *service) GetByPlant(ids []string) ([]domain.ListProject, error) {
	projects, err := s.repository.FindByPlant(ids)
	return projects, err
}

func (s *service) Store(input domain.ListProject) (domain.ListProject, error) {
	listProject, err := s.repository.Store(input)
	return listProject, err
}

func (s *service) Store2(input domain.ListProject2) (domain.ListProject2, error) {
	listProject, err := s.repository.Store2(input)
	return listProject, err
}

func (s *service) Store3(input domain.ListProject3) (domain.ListProject3, error) {
	listProject, err := s.repository.Store3(input)
	return listProject, err
}

func (s *service) Store4(input domain.ListProject4) (domain.ListProject4, error) {
	listProject, err := s.repository.Store4(input)
	return listProject, err
}

func (s *service) UpdateStatus(input domain.UpdateStatus, Id string) (domain.UpdateStatus, error) {
	projectList, err := s.repository.FindByPekerjaanNo(input.PekerjaanNo)
	if err != nil {
		return projectList, err
	}

	updateStatus := domain.UpdateStatus{
		PekerjaanNo: projectList.PekerjaanNo,
		Status:      input.Status,
	}

	status, err := s.repository.UpdateStatus(updateStatus)
	return status, err
}

func (s *service) UpdateSpkNRetensi(input domain.UpdateSpkNRetensi, Id string) (domain.UpdateSpkNRetensi, error) {
	status, err := s.repository.UpdateSpkNRetensi(input, Id)
	return status, err
}

func (s *service) UpdatePlanningActualDate(input domain.ModelUpdateActualPlanningDate) error {
	err := s.repository.UpdatePlanningActualDate(input)
	return err
}
