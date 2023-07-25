package boqheader

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error)
	Store(input domain.BoqHeader) (domain.BoqHeader, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error) {
	var boqHeaders []domain.BoqHeader
	err := r.db.Table("boq_header").Find(&boqHeaders).Error
	return boqHeaders, err
}

func (r *repository) Store(input domain.BoqHeader) (domain.BoqHeader, error) {

	err := r.db.Table("boq_header").Create(&input).Error
	return input, err
}
