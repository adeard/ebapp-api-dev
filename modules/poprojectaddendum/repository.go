package poprojectaddendum

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindByPo(po string) ([]domain.PoProjectAddendum, error)
	Delete(id string, po string, item string) error
	Store(input domain.PoProjectAddendum) (domain.PoProjectAddendum, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindByPo(po string) ([]domain.PoProjectAddendum, error) {
	var poProjectAddendum []domain.PoProjectAddendum

	q := r.db.Table("po_project_addendum").Debug()

	if po != "" {
		q = q.Where("pekerjaan_no = ?", po)
	}

	err := q.Order("[order] asc").Find(&poProjectAddendum).Error

	return poProjectAddendum, err
}

func (r *repository) Delete(id string, po string, item string) error {
	err := r.db.Table("po_project_addendum").Where("pekerjaan_no =?", id).Where("po =?", po).Where("[order] =?", item).Delete(&domain.PoProjectAddendum{}).Error
	return err
}

func (r *repository) Store(input domain.PoProjectAddendum) (domain.PoProjectAddendum, error) {
	err := r.db.Table("po_project_addendum").Create(&input).Error
	return input, err
}
