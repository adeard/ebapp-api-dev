package poprogressheader

import (
	"ebapp-api-dev/domain"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindProg(id string) (domain.PoProgressHeader, error)
	FindAllProg(id string) ([]domain.PoProgressHeaderWithPercentage, error)
	Delete(id string) error
	Update(id string, input domain.PoProgressHeaderUpdate) (domain.PoProgressHeader, error)
	EbappUpdate(id string, input domain.PoProgressHeaderUpdateEbapp) (domain.PoProgressHeader, error)
	EbappUpdate2(id string, status string) (domain.PoProgressHeader, error)
	Store(input domain.PoProgressHeader) (domain.PoProgressHeader, error)
	SyncProgress(run_num string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindProg(id string) (domain.PoProgressHeader, error) {
	var progress domain.PoProgressHeader
	err := r.db.Table("po_progress_header").Where("run_num =?", id).First(&progress).Error
	return progress, err
}

func (r *repository) FindAllProg(id string) ([]domain.PoProgressHeaderWithPercentage, error) {
	var progress []domain.PoProgressHeaderWithPercentage
	// err := r.db.Table("po_progress_header").Where("run_num LIKE ?", id+"%").Find(&progress).Error

	err := r.db.Table("po_progress_header").
		Select("po_progress_header.*, po_boq_header_progress.actual_percentage").
		Joins("JOIN po_boq_header_progress ON po_progress_header.run_num = po_boq_header_progress.pekerjaan_no").
		Where("po_progress_header.run_num LIKE ? AND po_boq_header_progress.[order] = 0", id+"%").
		Find(&progress).Error

	return progress, err
}

func (r *repository) Delete(id string) error {
	var progress domain.PoProgressHeader
	err := r.db.Table("po_progress_header").Where("run_num =?", id).Delete(&progress).Error
	return err
}

func (r *repository) Store(input domain.PoProgressHeader) (domain.PoProgressHeader, error) {
	err := r.db.Table("po_progress_header").Create(&input).Error
	return input, err
}

func (r *repository) Update(id string, input domain.PoProgressHeaderUpdate) (domain.PoProgressHeader, error) {
	updateData := map[string]interface{}{
		"new_prog": input.NewProg,
	}

	err := r.db.Table("po_progress_header").Where("run_num = ?", id).Updates(updateData).Error
	var data domain.PoProgressHeader
	if err != nil {
		return data, err
	}

	err = r.db.Table("po_progress_header").Where("run_num = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) EbappUpdate(id string, input domain.PoProgressHeaderUpdateEbapp) (domain.PoProgressHeader, error) {
	updateData := map[string]interface{}{
		"isebapp":      0,
		"last_updated": input.LastUpdated,
		"keterangan":   input.Keterangan,
	}

	if input.IsEbapp == 1 {
		updateData["isebapp"] = 1
	}

	err := r.db.Table("po_progress_header").Where("run_num = ?", id).Updates(updateData).Error
	var data domain.PoProgressHeader
	if err != nil {
		return data, err
	}

	err = r.db.Table("po_progress_header").Where("run_num = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) EbappUpdate2(id string, status string) (domain.PoProgressHeader, error) {
	updateData := map[string]interface{}{
		"status":       status,
		"last_updated": time.Now(),
	}

	err := r.db.Table("po_progress_header").Where("run_num = ?", id).Updates(updateData).Error
	var data domain.PoProgressHeader
	if err != nil {
		return data, err
	}

	err = r.db.Table("po_progress_header").Where("run_num = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) SyncProgress(run_num string) error {
	parts := strings.Split(run_num, "/")
	newStr := strings.Join(parts[:len(parts)-1], "/")
	query := `SELECT (SUM(CASE 
		WHEN (a.current_volume*a.price) is null then 0 
		ELSE (a.current_volume*a.price)
		END)/ SUM(a.price * a.qty))*100 as p
  FROM po_boq_body_progress a left join po_boq_header b on a.[order] = b.[order]  where a.run_num = ? and a.qty != 0 and b.is_addendum = 0 and b.pekerjaan_no = ?`
	var p float64
	r.db.Raw(query, run_num, newStr).First(&p)
	query2 := `UPDATE po_progress_header set new_prog = ?, last_updated = ? WHERE run_num = ?`
	err := r.db.Exec(query2, p, time.Now().UTC(), run_num)

	// if err == nil {
	query1 := `SELECT (SUM(CASE 
		WHEN (a.current_volume*a.price) is null then 0 
		ELSE (a.current_volume*a.price)
		END)/ SUM(a.price * a.qty))*100 as p
  FROM po_boq_body_progress a left join po_boq_header b on a.[order] = b.[order]  where a.run_num = ? and a.qty != 0 and b.is_addendum = 1 and b.pekerjaan_no = ?`
	var p1 float64
	r.db.Raw(query1, run_num, newStr).First(&p1)
	query21 := `UPDATE po_progress_header_addendum set new_prog = ?, last_updated = ? WHERE run_num = ?`
	r.db.Exec(query21, p1, time.Now().UTC(), run_num)

	// 	return err1.Error
	// }
	return err.Error
}
