package poprogressheader

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindProg(id string) (domain.PoProgressHeader, error)
	FindAllProg(id string) ([]domain.PoProgressHeader, error)
	Delete(id string) error
	Update(id string, input domain.PoProgressHeaderUpdate) (domain.PoProgressHeader, error)
	EbappUpdate(id string, input domain.PoProgressHeaderUpdateEbapp) (domain.PoProgressHeader, error)
	Store(input domain.PoProgressHeader) (domain.PoProgressHeader, error)
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
	err := r.db.Table("po_progress_header").Where("run_num = ?", id).Updates(input).Error
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

	err := r.db.Debug().Table("po_progress_header").Where("run_num = ?", id).Updates(updateData).Error
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
