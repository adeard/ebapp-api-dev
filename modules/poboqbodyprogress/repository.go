package poboqbodyprogress

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.PoBoqBodyProgress) (domain.PoBoqBodyProgress, error)
	Update(runNum string, order string, mainId int, parentId int, current_volume float64) (domain.PoBoqBodyProgress, error)
	FindByItemNo(itemNo string) (domain.PoBoqBodyProgress, error)
	FindByRunNum(runNum string, order string) ([]domain.PoBoqBodyProgress, error)
	CountRunNum(runNum string) (int, error)
	SelectMaxOrder(runNum string) (int, error)
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(input domain.PoBoqBodyProgress) (domain.PoBoqBodyProgress, error) {
	err := r.db.Table("po_boq_body_progress").Create(&input).Error
	return input, err
}

func (r *repository) FindByItemNo(itemNo string) (domain.PoBoqBodyProgress, error) {
	var poBoqBodyProgress domain.PoBoqBodyProgress
	err := r.db.Table("po_boq_body_progress").Where("item_no = ?", itemNo).First(&poBoqBodyProgress).Error
	return poBoqBodyProgress, err
}

func (r *repository) FindByRunNum(runNum string, order string) ([]domain.PoBoqBodyProgress, error) {
	var boqBody []domain.PoBoqBodyProgress

	q := r.db.Table("po_boq_body_progress")

	if runNum != "" {
		q = q.Where("run_num = ?", runNum).Where("[order] = ?", order)
	}

	err := q.Order("main_id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) CountRunNum(runNum string) (int, error) {
	var total int

	query := "SELECT COUNT(*) AS total FROM eBAPP.dbo.po_boq_body_progress WHERE run_num = ?"

	err := r.db.Raw(query, runNum).Scan(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *repository) SelectMaxOrder(runNum string) (int, error) {
	var total int

	query := "SELECT MAX(CAST([order] AS INT)) AS max_order	FROM eBAPP.dbo.po_boq_body_progress	WHERE run_num = ?"

	err := r.db.Raw(query, runNum).Scan(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *repository) Delete(id string) error {
	err := r.db.Table("po_boq_body_progress").Where("run_num = ?", id).Delete(&domain.PoBoqBodyProgress{}).Error
	return err
}

func (r *repository) Update(runNum string, order string, mainId int, parentId int, current_volume float64) (domain.PoBoqBodyProgress, error) {
	// Membuat variabel untuk menampung hasil pembaruan
	var updatedProgress domain.PoBoqBodyProgress

	// Menggunakan fungsi Update dari GORM untuk memperbarui data di database
	err := r.db.Table("po_boq_body_progress").
		Where("run_num = ? AND [order] = ? AND main_id = ? AND parent_id = ?", runNum, order, mainId, parentId).
		Updates(map[string]interface{}{"current_volume": current_volume}).
		Error

	if err != nil {
		// Mengembalikan error jika terjadi kesalahan saat pembaruan
		return updatedProgress, err
	}

	// Mengembalikan data yang telah diperbarui
	updatedProgress = domain.PoBoqBodyProgress{
		RunNum:        runNum,
		Order:         order,
		Id:            mainId,
		ParentId:      parentId,
		CurrentVolume: current_volume, // Menggunakan nilai baru untuk current_volume
	}

	return updatedProgress, nil
}
