package listproject

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.ListProjectRequest) ([]domain.ListProject, error)
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
