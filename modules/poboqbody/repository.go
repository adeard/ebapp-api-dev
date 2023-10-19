package poboqbody

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindByRunNum(runNum string, order string) ([]domain.PoBoqBody, error)
	Store(input domain.PoBoqBody) (domain.PoBoqBody, error)
	FindByItemNo(itemNo string) (domain.PoBoqBody, error)
	FindBoq(runNum string, order string, mainId string) ([]domain.PoBoqBody, error)
	Update(input domain.PoBoqBody) (domain.PoBoqBody, error)
	Delete(id string, orderId string, mainId string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByRunNum(runNum string, order string) ([]domain.PoBoqBody, error) {
	var boqBody []domain.PoBoqBody

	q := r.db.Table("po_boq_body")

	if runNum != "" {
		q = q.Where("run_num = ?", runNum).Where("[order] = ?", order)
	}

	err := q.Order("main_id asc").Find(&boqBody).Error

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

func (r *repository) FindBoq(runNum string, order string, mainId string) ([]domain.PoBoqBody, error) {
	var boqBody []domain.PoBoqBody

	q := r.db.Table("po_boq_body")

	if runNum != "" {
		q = q.Where("run_num = ?", runNum).Where("[order] = ?", order).Where("main_id = ?", mainId)
	}

	err := q.Order("main_id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) Delete(id string, orderId string, mainId string) error {
	err := r.db.Table("po_boq_body").Where("run_num =?", id).Where("[order] =?", orderId).Where("main_id =?", mainId).Delete(&domain.PoBoqBody{}).Error
	return err
}

func (r *repository) Update(input domain.PoBoqBody) (domain.PoBoqBody, error) {
	err := r.db.Table("po_boq_body").Where("run_num =?", input.RunNum).Where("[order] =?", input.Order).Where("main_id =?", input.Id).Save(&input).Error
	return input, err
}
