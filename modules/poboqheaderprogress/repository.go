package poboqheaderprogress

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindProgress(id string) ([]domain.PoBoqHeaderProgress, error)
	Store(input domain.PoBoqHeaderProgress) (domain.PoBoqHeaderProgress, error)
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProgress(id string) ([]domain.PoBoqHeaderProgress, error) {
	var headers []domain.PoBoqHeaderProgress

	q := r.db.Table("po_boq_header_progress")

	if id != "" {
		q = q.Where("pekerjaan_no = ?", id)
	}

	err := q.Order("'order' asc").Find(&headers).Error

	return headers, err
}

func (r *repository) Store(input domain.PoBoqHeaderProgress) (domain.PoBoqHeaderProgress, error) {
	err := r.db.Table("po_boq_header_progress").Create(&input).Error
	return input, err
}

func (r *repository) Delete(id string) error {
	err := r.db.Table("po_boq_header_progress").Where("pekerjaan_no =?", id).Delete(&domain.PoBoqHeaderProgress{}).Error
	return err
}
