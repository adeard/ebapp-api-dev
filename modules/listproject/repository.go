package listproject

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.ListProjectRequest) ([]domain.ListProject, error)
	FindById(id string) (domain.ListProject, error)
	Store(input domain.ListProject) (domain.ListProject, error)
	Store2(input domain.ListProject2) (domain.ListProject2, error)
	Store3(input domain.ListProject3) (domain.ListProject3, error)
	Store4(input domain.ListProject4) (domain.ListProject4, error)

	FindByPekerjaanNo(id string) (domain.UpdateStatus, error)
	UpdateStatus(input domain.UpdateStatus) (domain.UpdateStatus, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.ListProjectRequest) ([]domain.ListProject, error) {
	var listProjects []domain.ListProject
	err := r.db.Table("list_project").Order("pekerjaan_date ASC").Find(&listProjects).Error
	return listProjects, err
}

func (r *repository) FindById(id string) (domain.ListProject, error) {
	var project domain.ListProject
	err := r.db.Table("list_project").Where("id =?", id).First(&project).Error
	return project, err
}

func (r *repository) Store(input domain.ListProject) (domain.ListProject, error) {
	err := r.db.Table("list_project").Create(&input).Error
	return input, err
}

func (r *repository) Store2(input domain.ListProject2) (domain.ListProject2, error) {
	err := r.db.Table("list_project").Create(&input).Error
	return input, err
}

func (r *repository) Store3(input domain.ListProject3) (domain.ListProject3, error) {
	err := r.db.Table("list_project").Create(&input).Error
	return input, err
}

func (r *repository) Store4(input domain.ListProject4) (domain.ListProject4, error) {
	err := r.db.Table("list_project").Create(&input).Error
	return input, err
}

func (r *repository) FindByPekerjaanNo(id string) (domain.UpdateStatus, error) {
	var project domain.UpdateStatus
	err := r.db.Table("list_project").Where("pekerjaan_no =?", id).First(&project).Error
	return project, err
}

func (r *repository) UpdateStatus(input domain.UpdateStatus) (domain.UpdateStatus, error) {
	err := r.db.Table("list_project").Where("pekerjaan_no =?", input.PekerjaanNo).Save(&input).Error
	return input, err
}
