package poprojectaddendum

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindByPo(po string) ([]domain.PoProjectAddendum, error)
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

	err := q.Order("id asc").Find(&poProjectAddendum).Error

	return poProjectAddendum, err
}
