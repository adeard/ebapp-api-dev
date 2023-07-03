package boqbody

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.BoqBodyRequest) ([]domain.BoqBody, error)
	Store(input domain.BoqBodyRequest) (domain.BoqBodyRequest, error)
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

	err := q.Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) Store(input domain.BoqBodyRequest) (domain.BoqBodyRequest, error) {
	err := r.db.Table("boq_body").Debug().Create(&input).Error

	return input, err
}
