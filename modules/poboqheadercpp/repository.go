package poboqheadercpp

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindCpp(id string) ([]domain.PoBoqHeaderCpp, error)
	Store(input domain.PoBoqHeaderCpp) (domain.PoBoqHeaderCpp, error)
	Delete(id string) error
	CloneProgress(oldIdProgress string, newIdCpp string) error
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
