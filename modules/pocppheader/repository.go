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
	CloneProgress(oldIdProgress string, newIdCpp string) error
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

// Clone CPP body from header
func (r *repository) CloneProgress(oldIdProgress string, newIdCpp string) error {
	// Lakukan query untuk meng-INSERT data baru berdasarkan data yang ada dengan menggunakan parameter oldId dan newId
	query := `INSERT INTO eBAPP.dbo.po_boq_body_cpp ([run_num]
		,[item_no]
		,[item_level]
		,[item_description]
		,[item_specification]
		,[qty]
		,[status]
		,[note]
		,[main_id]
		,[parent_id]
		,[order])
               SELECT 
                   ? AS run_num,
                   item_no,
                   item_level,
                   item_description,
                   item_specification,
                   qty,
                   0,
                   null,
				   main_id,
                   parent_id,
				   [order]
               FROM eBAPP.dbo.po_boq_body_progress
               WHERE run_num = ?`

	// Eksekusi query dengan parameter oldId dan newId
	result := r.db.Exec(query, newIdCpp, oldIdProgress)

	// Periksa jika terjadi kesalahan saat menjalankan query
	if result.Error != nil {
		return result.Error
	}

	// Jika berhasil, kembalikan pesan berhasil
	return nil
}
