package poboqheadercpp

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindCpp(id string) ([]domain.PoBoqHeaderCpp, error)
	Store(input domain.PoBoqHeaderCpp) (domain.PoBoqHeaderCpp, error)
	Delete(id string) error
	FindCppWithPagingServerSide(id string, page int, pageSize int, item string, desc string) ([]domain.PoBoqHeaderCpp, error)
	FindCppWithPagingServerSideCount(id string, item string, desc string) (int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCpp(id string) ([]domain.PoBoqHeaderCpp, error) {
	var headers []domain.PoBoqHeaderCpp

	q := r.db.Table("po_boq_header_cpp")

	if id != "" {
		q = q.Where("pekerjaan_no = ?", id)
	}

	err := q.Order("'order' asc").Find(&headers).Error

	return headers, err
}

func (r *repository) Store(input domain.PoBoqHeaderCpp) (domain.PoBoqHeaderCpp, error) {
	err := r.db.Table("po_boq_header_cpp").Create(&input).Error
	return input, err
}

func (r *repository) Delete(id string) error {
	err := r.db.Table("po_boq_header_cpp").Where("pekerjaan_no =?", id).Delete(&domain.PoBoqHeaderCpp{}).Error
	return err
}

func (r *repository) FindCppWithPagingServerSide(id string, page int, pageSize int, item string, desc string) ([]domain.PoBoqHeaderCpp, error) {
	var headers []domain.PoBoqHeaderCpp

	q := r.db.Table("po_boq_header_cpp")

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

func (r *repository) FindCppWithPagingServerSideCount(id string, item string, desc string) (int64, error) {
	var count int64
	q := r.db.Table("po_boq_header_cpp")

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
