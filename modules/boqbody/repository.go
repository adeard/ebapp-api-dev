package boqbody

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.BoqBodyRequest) ([]domain.BoqBody, error)
	FindByRunNum(runNum string) ([]domain.BoqBody, error)
	FindByParentID(parentID string) ([]domain.BoqBody, error)
	FindById(id string) (domain.BoqBody, error)
	FindByItemNo(itemNo string) (domain.BoqBody, error)
	Store(input domain.BoqBody) (domain.BoqBody, error)
	Update(input domain.BoqBody) (domain.BoqBody, error)
	DeleteByID(id string) error
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

func (r *repository) FindByRunNum(runNum string) ([]domain.BoqBody, error) {
	var boqBody []domain.BoqBody

	q := r.db.Table("boq_body")

	if runNum != "" {
		q = q.Where("run_num = ?", runNum)
	}

	err := q.Order("id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) FindByParentID(parentID string) ([]domain.BoqBody, error) {
	var boqBody []domain.BoqBody

	q := r.db.Table("boq_body").Debug()

	if parentID != "" {
		q = q.Where("parent_id = ?", parentID)
	}

	err := q.Order("parent_id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) FindById(id string) (domain.BoqBody, error) {
	var boqBody domain.BoqBody
	err := r.db.Table("boq_body").Where("id =?", id).First(&boqBody).Error
	return boqBody, err
}

func (r *repository) FindByItemNo(itemNo string) (domain.BoqBody, error) {
	var boqBody domain.BoqBody
	err := r.db.Table("boq_body").Where("item_no = ?", itemNo).First(&boqBody).Error
	return boqBody, err
}

func (r *repository) Store(input domain.BoqBody) (domain.BoqBody, error) {
	query := `SELECT TOP 1 id FROM boq_body order by id desc`
	var id int
	r.db.Raw(query).First(&id)
	query2 := `INSERT INTO boq_body([run_num]
		,[item_no]
		,[item_level]
		,[item_description]
		,[item_specification]
		,[qty]
		,[unit]
		,[price]
		,[currency]
		,[note]
		,[id]
		,[parent_id]) values(?,?,?,?,?,?,?,?,?,?,?,?)`
	err := r.db.Exec(query2, input.RunNum, input.ItemNo, input.ItemLevel, input.ItemDescription, input.ItemSpecification, input.Qty, input.Unit, input.Price, input.Currency, input.Note, id+1, input.ParentId)
	return input, err.Error
}

func (r *repository) Update(input domain.BoqBody) (domain.BoqBody, error) {
	err := r.db.Table("boq_body").Where("id =?", input.Id).Save(&input).Error
	return input, err
}

func (r *repository) DeleteByID(id string) error {
	err := r.db.Table("boq_body").Where("id =?", id).Delete(&domain.BoqBody{}).Error
	return err
}
