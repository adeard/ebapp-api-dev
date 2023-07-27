package boqheader

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error)
	FindById(id string) (domain.BoqHeader, error)
	Store(input domain.BoqHeader) (domain.BoqHeader, error)
	Update(input domain.BoqHeader) (domain.BoqHeader, error)
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

func (r *repository) FindById(id string) (domain.BoqHeader, error) {
	var boqHeader domain.BoqHeader
	err := r.db.Table("boq_header").Where("run_num =?", id).First(&boqHeader).Error
	return boqHeader, err
}

func (r *repository) Store(input domain.BoqHeader) (domain.BoqHeader, error) {

	err := r.db.Table("boq_header").Create(&input).Error
	return input, err
}

func (r *repository) Update(input domain.BoqHeader) (domain.BoqHeader, error) {
	err := r.db.Table("boq_header").Where("run_num =?", input.RunNum).Save(&input).Error
	return input, err
}
