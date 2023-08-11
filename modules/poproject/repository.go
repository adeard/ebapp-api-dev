package poproject

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.PoProjectRequest) ([]domain.PoProject, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.PoProjectRequest) ([]domain.PoProject, error) {
	var poProject []domain.PoProject
	err := r.db.Table("po_project").Find(&poProject).Error
	return poProject, err
}
