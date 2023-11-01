package poboqheader

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindByPekerjaanNo(id string) ([]domain.PoBoqHeader, error)
	Delete(id string, po string, order string) error
	Store(input domain.PoBoqHeader) (domain.PoBoqHeader, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByPekerjaanNo(id string) ([]domain.PoBoqHeader, error) {
	var headers []domain.PoBoqHeader

	q := r.db.Table("po_boq_header")

	if id != "" {
		q = q.Where("pekerjaan_no = ?", id)
	}

	err := q.Order("'order' asc").Find(&headers).Error

	return headers, err
}

func (r *repository) Delete(id string, po string, item string) error {
	err := r.db.Table("po_boq_header").Where("pekerjaan_no =?", id).Where("po =?", po).Where("[order] =?", item).Where("is_addendum = 1").Delete(&domain.PoBoqHeader{}).Error
	return err
}

func (r *repository) Store(input domain.PoBoqHeader) (domain.PoBoqHeader, error) {
	err := r.db.Table("po_boq_header").Create(&input).Error
	return input, err
}
