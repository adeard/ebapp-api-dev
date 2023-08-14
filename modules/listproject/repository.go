package listproject

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.ListProjectRequest) ([]domain.ListProject, error)
	FindById(id string) (domain.ListProject, error)
	Store(input domain.ListProject) (domain.ListProject, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.ListProjectRequest) ([]domain.ListProject, error) {
	var listProjects []domain.ListProject
	err := r.db.Table("list_project").Find(&listProjects).Error
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
