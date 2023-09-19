package poproject

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.PoProjectRequest) ([]domain.PoProject, error)
	FindByPo(po string, no string) ([]domain.PoProject, error)
	Store(input domain.PoProject) (domain.PoProject, error)
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

func (r *repository) FindByPo(po string, no string) ([]domain.PoProject, error) {
	var poProject []domain.PoProject

	q := r.db.Table("po_project").Debug()

	if po != "" {
		q = q.Where("po = ?", po).Where("pekerjaan_no = ?", no)
	}

	err := q.Order("id asc").Find(&poProject).Error

	return poProject, err
}

func (r *repository) Store(input domain.PoProject) (domain.PoProject, error) {
	err := r.db.Table("po_project").Create(&input).Error
	return input, err
}
