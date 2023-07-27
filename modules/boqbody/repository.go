package boqbody

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.BoqBodyRequest) ([]domain.BoqBody, error)
	FindByID(runNum string) ([]domain.BoqBody, error)
	Store(input domain.BoqBody) (domain.BoqBody, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.BoqBodyRequest) ([]domain.BoqBody, error) {
	var boqBody []domain.BoqBody

	q := r.db.Table("boq_body").Debug()

	if input.RunNum != "" {
		q = q.Where("run_num = ?", input.RunNum)
	}

	err := q.Order("id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) FindByID(runNum string) ([]domain.BoqBody, error) {
	var boqBody []domain.BoqBody

	q := r.db.Table("boq_body").Debug()

	if runNum != "" {
		q = q.Where("run_num = ?", runNum)
	}

	err := q.Order("id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) Store(input domain.BoqBody) (domain.BoqBody, error) {
	err := r.db.Table("boq_body").Create(&input).Error
	return input, err
}
