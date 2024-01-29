package poboqbodyprogress

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.PoBoqBodyProgress) (domain.PoBoqBodyProgress, error)
	FindByItemNo(itemNo string) (domain.PoBoqBodyProgress, error)
	FindByRunNum(runNum string, order string) ([]domain.PoBoqBodyProgress, error)
	Delete(id string) error
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

func (r *repository) FindByRunNum(runNum string, order string) ([]domain.PoBoqBodyProgress, error) {
	var boqBody []domain.PoBoqBodyProgress

	q := r.db.Table("po_boq_body_progress")

	if runNum != "" {
		q = q.Where("run_num = ?", runNum).Where("[order] = ?", order)
	}

	err := q.Order("main_id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) Delete(id string) error {
	err := r.db.Table("po_boq_body_progress").Where("run_num = ?", id).Delete(&domain.PoBoqBodyProgress{}).Error
	return err
}
