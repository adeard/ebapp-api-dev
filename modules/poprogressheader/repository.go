package poprogressheader

import (
	"ebapp-api-dev/domain"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindProg(id string) (domain.PoProgressHeader, error)
	FindAllProg(id string) ([]domain.PoProgressHeader, error)
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

func (r *repository) FindAllProg(id string) ([]domain.PoProgressHeader, error) {
	var progress []domain.PoProgressHeader
	err := r.db.Table("po_progress_header").Where("run_num LIKE ?", id+"%").Find(&progress).Error
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
	query := `SELECT (SUM(CASE 
		WHEN (current_volume*price) is null then 0 
		ELSE (current_volume*price)
		END)/ SUM(price * qty))*100 as p
  FROM po_boq_body_progress where run_num = ? and qty != 0`
	var p float64
	r.db.Raw(query, run_num).First(&p)
	query2 := `UPDATE po_progress_header set new_prog = ?, last_updated = ? WHERE run_num = ?`
	err := r.db.Exec(query2, p, time.Now().UTC(), run_num)

	return err.Error
}
