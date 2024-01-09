package poboqbodyprogress

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.PoBoqBodyProgress) (domain.PoBoqBodyProgress, error)
	FindByItemNo(itemNo string) (domain.PoBoqBodyProgress, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(input domain.PoBoqBodyProgress) (domain.PoBoqBodyProgress, error) {
	err := r.db.Table("po_boq_body_progress").Create(&input).Error
	return input, err
}

func (r *repository) FindByItemNo(itemNo string) (domain.PoBoqBodyProgress, error) {
	var poBoqBodyProgress domain.PoBoqBodyProgress
	err := r.db.Table("po_boq_body_progress").Where("item_no = ?", itemNo).First(&poBoqBodyProgress).Error
	return poBoqBodyProgress, err
}
