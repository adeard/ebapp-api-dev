package boqheader

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error)
	FindAllActive(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error)
	FindById(id string) (domain.BoqHeader, error)
	Store(input domain.BoqHeader) (domain.BoqHeader, error)
	Update(input domain.BoqHeader) (domain.BoqHeader, error)
	Clone(oldId string, newId string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error) {
	var boqHeaders []domain.BoqHeader
	err := r.db.Table("boq_header").Find(&boqHeaders).Error
	return boqHeaders, err
}

func (r *repository) FindAllActive(input domain.BoqHeaderRequest) ([]domain.BoqHeader, error) {
	var boqHeaders []domain.BoqHeader
	err := r.db.Table("boq_header").Where("header_status =?", true).Find(&boqHeaders).Error
	return boqHeaders, err
}

func (r *repository) FindById(id string) (domain.BoqHeader, error) {
	var boqHeader domain.BoqHeader
	err := r.db.Table("boq_header").Where("run_num =?", id).First(&boqHeader).Error
	return boqHeader, err
}

func (r *repository) Store(input domain.BoqHeader) (domain.BoqHeader, error) {
	err := r.db.Table("boq_header").Create(&input).Error
	return input, err
}

func (r *repository) Update(input domain.BoqHeader) (domain.BoqHeader, error) {
	err := r.db.Table("boq_header").Where("run_num =?", input.RunNum).Save(&input).Error
	return input, err
}

// Clone body from header
func (r *repository) Clone(oldId string, newId string) (string, error) {
	// Lakukan query untuk meng-INSERT data baru berdasarkan data yang ada dengan menggunakan parameter oldId dan newId
	query := `INSERT INTO eBAPP.dbo.boq_body (run_num, item_no, item_level, item_description, item_specification, qty, unit, price, currency, note, parent_id)
               SELECT 
                   ? AS run_num,
                   item_no,
                   item_level,
                   item_description,
                   item_specification,
                   qty,
                   unit,
                   price,
                   currency,
                   note,
                   parent_id
               FROM eBAPP.dbo.boq_body
               WHERE run_num = ?`

	// Eksekusi query dengan parameter oldId dan newId
	result := r.db.Exec(query, newId, oldId)

	// Periksa jika terjadi kesalahan saat menjalankan query
	if result.Error != nil {
		return "", result.Error
	}

	// Jika berhasil, kembalikan pesan berhasil
	message := "Data BOQ berhasil diduplikasi"
	return message, nil
}
