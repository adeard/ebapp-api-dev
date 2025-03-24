package poboqheaderprogress

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindProgress(id string) ([]domain.PoBoqHeaderProgress, error)
	Store(input domain.PoBoqHeaderProgress) (domain.PoBoqHeaderProgress, error)
	Delete(id string) error
	FindProgressWithPagingServerSide(id string, page int, pageSize int, item string, desc string) ([]domain.PoBoqHeaderProgress, error)
	FindProgressWithPagingServerSideCount(id string, item string, desc string) (int64, error)
	SyncProgress(run_num string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProgress(id string) ([]domain.PoBoqHeaderProgress, error) {
	var headers []domain.PoBoqHeaderProgress

	q := r.db.Table("po_boq_header_progress")

	if id != "" {
		q = q.Where("pekerjaan_no = ?", id)
	}

	err := q.Order("'order' asc").Find(&headers).Error

	return headers, err
}

func (r *repository) Store(input domain.PoBoqHeaderProgress) (domain.PoBoqHeaderProgress, error) {
	err := r.db.Table("po_boq_header_progress").Create(&input).Error
	return input, err
}

func (r *repository) Delete(id string) error {
	err := r.db.Table("po_boq_header_progress").Where("pekerjaan_no =?", id).Delete(&domain.PoBoqHeaderProgress{}).Error
	return err
}

func (r *repository) FindProgressWithPagingServerSide(id string, page int, pageSize int, item string, desc string) ([]domain.PoBoqHeaderProgress, error) {
	var headers []domain.PoBoqHeaderProgress

	q := r.db.Table("po_boq_header_progress")

	if item == "" && desc == "" {
		query := `pekerjaan_no = ?`
		q = q.Where(query, id)
	} else if item != "" && desc == "" {
		query := `pekerjaan_no = ? and item = ?`
		q = q.Where(query, id, item)
	} else if item == "" && desc != "" {
		query := `pekerjaan_no = ? and description LIKE ?`
		q = q.Where(query, id, "%"+desc+"%")
	} else {
		query := `pekerjaan_no = ? and item = ? and description LIKE ?`
		q = q.Where(query, id, item, "%"+desc+"%")
	}

	err := q.
		Order("CAST([order] AS INT) ASC").
		Limit(pageSize).
		Offset(pageSize * (page - 1)).
		Find(&headers).
		Error

	return headers, err
}

func (r *repository) FindProgressWithPagingServerSideCount(id string, item string, desc string) (int64, error) {
	var count int64
	q := r.db.Table("po_boq_header_progress")

	if item == "" && desc == "" {
		query := `pekerjaan_no = ?`
		q = q.Where(query, id)
	} else if item != "" && desc == "" {
		query := `pekerjaan_no = ? and item = ?`
		q = q.Where(query, id, item)
	} else if item == "" && desc != "" {
		query := `pekerjaan_no = ? and description LIKE ?`
		q = q.Where(query, id, "%"+desc+"%")
	} else {
		query := `pekerjaan_no = ? and item = ? and description LIKE ?`
		q = q.Where(query, id, item, "%"+desc+"%")
	}

	err := q.Count(&count).Error
	return count, err
}

func (r *repository) SyncProgress(run_num string) error {
	// headers, err := r.FindProgress(run_num)

	// for _, headersData := range headers {
	query := `SELECT (SUM(CASE 
		WHEN (current_volume*price) is null then 0 
		ELSE (current_volume*price)
		END + 
		CASE 
		WHEN (previous_volume*price) is null then 0 
		ELSE (previous_volume*price)
		END)/ SUM(price * qty))*100 as p
  FROM po_boq_body_progress where run_num = ? and qty != 0`
	var p float64
	r.db.Raw(query, run_num).First(&p)
	query2 := `UPDATE po_boq_header_progress set actual_percentage = ? WHERE pekerjaan_no = ? and [order] = ?`
	err2 := r.db.Exec(query2, p, run_num, 0)
	// if err2 != nil {
	// 	return err2.Error
	// }

	//}

	return err2.Error
}
