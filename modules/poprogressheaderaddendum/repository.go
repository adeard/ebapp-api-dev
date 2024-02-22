package poprogressheaderaddendum

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAllProg(id string) ([]domain.PoProgressHeaderAddendum, error)
	Store(input domain.PoProgressHeaderAddendum) (domain.PoProgressHeaderAddendum, error)
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAllProg(id string) ([]domain.PoProgressHeaderAddendum, error) {
	var progress []domain.PoProgressHeaderAddendum
	err := r.db.Table("po_progress_header_addendum").Where("run_num =?", id).Find(&progress).Error
	return progress, err
}

func (r *repository) Store(input domain.PoProgressHeaderAddendum) (domain.PoProgressHeaderAddendum, error) {
	err := r.db.Table("po_progress_header_addendum").Create(&input).Error
	return input, err
}

func (r *repository) Delete(id string) error {
	var progress domain.PoProgressHeaderAddendum
	err := r.db.Table("po_progress_header_addendum").Where("run_num =?", id).Delete(&progress).Error
	return err
}
