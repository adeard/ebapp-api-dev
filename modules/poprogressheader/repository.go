package poprogressheader

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindProg(id string) (domain.PoProgressHeader, error)
	FindAllProg(id string) ([]domain.PoProgressHeader, error)
	Store(input domain.PoProgressHeader) (domain.PoProgressHeader, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindProg(id string) (domain.PoProgressHeader, error) {
	var progress domain.PoProgressHeader
	err := r.db.Table("po_progress_header").Where("run_num =?", id).First(&progress).Error
	return progress, err
}

func (r *repository) FindAllProg(id string) ([]domain.PoProgressHeader, error) {
	var progress []domain.PoProgressHeader
	err := r.db.Table("po_progress_header").Where("run_num LIKE ?", id+"%").Find(&progress).Error
	return progress, err
}

func (r *repository) Store(input domain.PoProgressHeader) (domain.PoProgressHeader, error) {
	err := r.db.Table("po_progress_header").Create(&input).Error
	return input, err
}
