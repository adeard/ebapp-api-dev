package poproject

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.PoProjectRequest) ([]domain.PoProject, error)
	FindByPo(po string, no string) ([]domain.PoProject, error)
	Store(input domain.PoProject) (domain.PoProject, error)

	FindCompanyByBA(ba string) ([]domain.Company, error)
	FindPlantByWreks(id string) ([]domain.Plant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindCompanyByBA(ba string) ([]domain.Company, error) {
	var addon []domain.Company

	q := r.db.Table("MasterCompany").Debug()

	if ba != "" {
		q = q.Where("BUKRS = ?", ba)
	}

	err := q.Find(&addon).Error

	return addon, err
}

func (r *repository) FindPlantByWreks(id string) ([]domain.Plant, error) {
	var addon []domain.Plant

	q := r.db.Table("MasterPlant").Debug()

	if id != "" {
		q = q.Where("WERKS = ?", id)
	}

	err := q.Find(&addon).Error

	return addon, err
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
