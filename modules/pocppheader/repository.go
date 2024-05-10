package pocppheader

import (
	"ebapp-api-dev/domain"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindCpp(id string) (domain.PoCppHeader, error)
	FindCppByRunNumProgress(id string) (domain.PoCppHeader, error)
	FindAllProg(id string) ([]domain.PoCppHeader, error)
	Delete(id string) error
	Update(id string, input domain.PoCppHeaderUpdate) (domain.PoCppHeader, error)
	UpdateStatus(id string, status string) (domain.PoCppHeader, error)
	Store(input domain.PoCppHeader) (domain.PoCppHeader, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindCpp(id string) (domain.PoCppHeader, error) {
	var Cpp domain.PoCppHeader
	err := r.db.Table("po_cpp_header").Where("run_num =?", id).First(&Cpp).Error
	return Cpp, err
}

func (r *repository) FindCppByRunNumProgress(id string) (domain.PoCppHeader, error) {
	var Cpp domain.PoCppHeader
	err := r.db.Table("po_cpp_header").Where("run_num_progress =?", id).First(&Cpp).Error
	return Cpp, err
}

func (r *repository) FindAllProg(id string) ([]domain.PoCppHeader, error) {
	var Cpp []domain.PoCppHeader
	err := r.db.Table("po_cpp_header").Where("run_num LIKE ?", id+"%").Find(&Cpp).Error
	return Cpp, err
}

func (r *repository) Delete(id string) error {
	var Cpp domain.PoCppHeader
	err := r.db.Table("po_cpp_header").Where("run_num =?", id).Delete(&Cpp).Error
	return err
}

func (r *repository) Store(input domain.PoCppHeader) (domain.PoCppHeader, error) {
	err := r.db.Table("po_cpp_header").Create(&input).Error
	return input, err
}

func (r *repository) Update(id string, input domain.PoCppHeaderUpdate) (domain.PoCppHeader, error) {
	err := r.db.Table("po_cpp_header").Where("run_num = ?", id).Updates(input).Error
	var data domain.PoCppHeader
	if err != nil {
		return data, err
	}

	err = r.db.Table("po_cpp_header").Where("run_num = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) UpdateStatus(id string, status string) (domain.PoCppHeader, error) {
	updateData := map[string]interface{}{
		"status_progress": status,
		"last_updated":    time.Now(),
	}

	err := r.db.Table("po_cpp_header").Where("run_num = ?", id).Updates(updateData).Error
	var data domain.PoCppHeader
	if err != nil {
		return data, err
	}

	err = r.db.Table("po_cpp_header").Where("run_num = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
