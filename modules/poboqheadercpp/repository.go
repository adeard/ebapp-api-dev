package poboqheadercpp

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindCpp(id string) ([]domain.PoBoqHeaderCpp, error)
	Store(input domain.PoBoqHeaderCpp) (domain.PoBoqHeaderCpp, error)
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCpp(id string) ([]domain.PoBoqHeaderCpp, error) {
	var headers []domain.PoBoqHeaderCpp

	q := r.db.Table("po_boq_header_cpp")

	if id != "" {
		q = q.Where("pekerjaan_no = ?", id)
	}

	err := q.Order("'order' asc").Find(&headers).Error

	return headers, err
}

func (r *repository) Store(input domain.PoBoqHeaderCpp) (domain.PoBoqHeaderCpp, error) {
	err := r.db.Table("po_boq_header_cpp").Create(&input).Error
	return input, err
}

func (r *repository) Delete(id string) error {
	err := r.db.Table("po_boq_header_cpp").Where("pekerjaan_no =?", id).Delete(&domain.PoBoqHeaderCpp{}).Error
	return err
}
