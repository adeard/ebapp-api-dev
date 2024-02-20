package poprogressheaderaddendum

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Delete(id string) error {
	var progress domain.PoProgressHeaderAddendum
	err := r.db.Table("po_progress_header_addendum").Where("run_num =?", id).Delete(&progress).Error
	return err
}
