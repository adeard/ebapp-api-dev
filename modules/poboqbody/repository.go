package poboqbody

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindByRunNum(runNum string) ([]domain.PoBoqBody, error)
	Store(input domain.PoBoqBody) (domain.PoBoqBody, error)
	FindByItemNo(itemNo string) (domain.PoBoqBody, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByRunNum(runNum string) ([]domain.PoBoqBody, error) {
	var boqBody []domain.PoBoqBody

	q := r.db.Table("po_boq_body")

	if runNum != "" {
		q = q.Where("run_num = ?", runNum)
	}

	err := q.Order("id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) Store(input domain.PoBoqBody) (domain.PoBoqBody, error) {
	err := r.db.Table("po_boq_body").Create(&input).Error
	return input, err
}

func (r *repository) FindByItemNo(itemNo string) (domain.PoBoqBody, error) {
	var poBoqBody domain.PoBoqBody
	err := r.db.Table("po_boq_body").Where("item_no = ?", itemNo).First(&poBoqBody).Error
	return poBoqBody, err
}
