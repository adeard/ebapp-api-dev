package boqbody

import (
	"ebapp-api-dev/domain"

	"strconv"

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
	DeleteByID(id string, runNum string) error

	FindByRunNumServerSide(runNum string, page int, pageSize int, item_no string, item_desc string) ([]domain.BoqBody, error)
	FindByRunNumServerSideCount(runNum string, item_no string, item_desc string) (int64, error)
	FindByParentIDServerSide(parentID string) ([]domain.BoqBody, error)
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

func (r *repository) FindByRunNumServerSide(runNum string, page int, pageSize int, item_no string, item_desc string) ([]domain.BoqBody, error) {
	var boqBody []domain.BoqBody

	q := r.db.Table("boq_body")

	if item_no == "" && item_desc == "" {
		query := `run_num = ? and item_level = 1`
		q = q.Where(query, runNum)
	} else if item_no != "" && item_desc == "" {
		query := `run_num = ? and item_no = ? and item_level = 1`
		q = q.Where(query, runNum, item_no)
	} else if item_no == "" && item_desc != "" {
		query := `run_num = ? and item_description LIKE ? and item_level = 1`
		q = q.Where(query, runNum, "%"+item_desc+"%")
	} else {
		query := `run_num = ? and item_no = ? and item_description LIKE ? and item_level = 1`
		q = q.Where(query, runNum, item_no, "%"+item_desc+"%")
	}

	err := q.
		Order("item_no ASC").
		Limit(pageSize).
		Offset(pageSize * (page - 1)).
		Find(&boqBody).
		Error

	return boqBody, err
}

func (r *repository) FindByRunNumServerSideCount(runNum string, item_no string, item_desc string) (int64, error) {
	var count int64

	q := r.db.Table("boq_body")

	if item_no == "" && item_desc == "" {
		query := `run_num = ? and item_level = 1`
		q = q.Where(query, runNum)
	} else if item_no != "" && item_desc == "" {
		query := `run_num = ? and item_no = ? and item_level = 1`
		q = q.Where(query, runNum, item_no)
	} else if item_no == "" && item_desc != "" {
		query := `run_num = ? and item_description LIKE ? and item_level = 1`
		q = q.Where(query, runNum, "%"+item_desc+"%")
	} else {
		query := `run_num = ? and item_no = ? and item_description LIKE ? and item_level = 1`
		q = q.Where(query, runNum, item_no, "%"+item_desc+"%")
	}

	err := q.Count(&count).Error
	return count, err
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

func (r *repository) FindByParentIDServerSide(parentID string) ([]domain.BoqBody, error) {
	var boqBody []domain.BoqBody

	q := r.db.Table("boq_body").Debug()

	if parentID != "" {
		q = q.Where("parent_id = ?", parentID)
	}

	err := q.Order("item_no asc").Find(&boqBody).Error

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
	input.Id = id + 1
	if input.ParentId == 0 {
		input.Order = strconv.Itoa(input.Id)
	} else {
		input.Order = strconv.Itoa(input.ParentId) + "-" + strconv.Itoa(input.Id)
	}

	return input, err.Error
}

func (r *repository) Update(input domain.BoqBody) (domain.BoqBody, error) {
	var parentId *int
	err := r.db.Table("boq_body").Select("parent_id").Where("run_num = ?", input.RunNum).Where("id = ?", input.Id).Scan(&parentId).Error
	if err != nil {
		return input, err
	}

	// Jika parentId di database adalah NULL, ubah menjadi 0
	if parentId == nil {
		err = r.db.Table("boq_body").Where("run_num = ?", input.RunNum).Where("id = ?", input.Id).Update("parent_id", 0).Error
		if err != nil {
			return input, err
		}
	}

	updateFields := map[string]interface{}{
		"item_no":            input.ItemNo,
		"item_description":   input.ItemDescription,
		"item_specification": input.ItemSpecification,
		"qty":                input.Qty,
		"unit":               input.Unit,
		"price":              input.Price,
		"currency":           input.Currency,
		"note":               input.Note,
	}

	err = r.db.Table("boq_body").Where("run_num = ?", input.RunNum).Where("id = ?", input.Id).Where("parent_id = ?", input.ParentId).Updates(updateFields).Error
	return input, err
}

func (r *repository) DeleteByID(id string, runNum string) error {
	err := r.db.Table("boq_body").Where("id =?", id).Where("run_num =?", runNum).Delete(&domain.BoqBody{}).Error
	return err
}
