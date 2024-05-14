package poboqbodycpp

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.PoBoqBodyCpp) (domain.PoBoqBodyCpp, error)
	Update(runNum string, order string, mainId int, parentId int, status bool, note string) (domain.PoBoqBodyCpp, error)
	FindByItemNo(itemNo string) (domain.PoBoqBodyCpp, error)
	FindByRunNum(runNum string, runNumProgress string, order string) ([]domain.PoBoqBodyCppProgress, error)
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

func (r *repository) Store(input domain.PoBoqBodyCpp) (domain.PoBoqBodyCpp, error) {
	err := r.db.Table("po_boq_body_cpp").Create(&input).Error
	return input, err
}

func (r *repository) FindByItemNo(itemNo string) (domain.PoBoqBodyCpp, error) {
	var poBoqBodyCpp domain.PoBoqBodyCpp
	err := r.db.Table("po_boq_body_cpp").Where("item_no = ?", itemNo).First(&poBoqBodyCpp).Error
	return poBoqBodyCpp, err
}

func (r *repository) FindByRunNum(runNum string, runNumProgress string, order string) ([]domain.PoBoqBodyCppProgress, error) {
	var boqBody []domain.PoBoqBodyCppProgress

	query := `SELECT DISTINCT a.*, b.status as status_cpp, b.note as note_cpp, b.run_num as run_num_cpp FROM po_boq_body_progress a right join po_boq_body_cpp b on a.item_no = b.item_no and a.item_level = b.item_level 
	and a.main_id = b.main_id and a.parent_id = b.parent_id where  b.run_num = ? and a.run_num = ? and a.[order] = ? order by a.main_id asc`
	err := r.db.Raw(query, runNum, runNumProgress, order).Find(&boqBody).Error

	// q := r.db.Table("po_boq_body_cpp")

	// if runNum != "" {
	// 	q = q.Where("run_num = ?", runNum).Where("[order] = ?", order)
	// }

	// err := q.Order("main_id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) CountRunNum(runNum string) (int, error) {
	var total int

	query := "SELECT COUNT(*) AS total FROM eBAPP.dbo.po_boq_body_cpp WHERE run_num = ?"

	err := r.db.Raw(query, runNum).Scan(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *repository) SelectMaxOrder(runNum string) (int, error) {
	var total int

	query := "SELECT MAX(CAST([order] AS INT)) AS max_order	FROM eBAPP.dbo.po_boq_body_cpp	WHERE run_num = ?"

	err := r.db.Raw(query, runNum).Scan(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *repository) Delete(id string) error {
	err := r.db.Table("po_boq_body_cpp").Where("run_num = ?", id).Delete(&domain.PoBoqBodyCpp{}).Error
	return err
}

func (r *repository) Update(runNum string, order string, mainId int, parentId int, status bool, note string) (domain.PoBoqBodyCpp, error) {
	// Membuat variabel untuk menampung hasil pembaruan
	var updatedCpp domain.PoBoqBodyCpp

	// Menggunakan fungsi Update dari GORM untuk memperbarui data di database
	err := r.db.Table("po_boq_body_cpp").
		Where("run_num = ? AND [order] = ? AND main_id = ? AND parent_id = ?", runNum, order, mainId, parentId).
		Updates(map[string]interface{}{"status": status, "note": note}).
		Error

	if err != nil {
		// Mengembalikan error jika terjadi kesalahan saat pembaruan
		return updatedCpp, err
	}

	// Mengembalikan data yang telah diperbarui
	updatedCpp = domain.PoBoqBodyCpp{
		RunNum:   runNum,
		Order:    order,
		Id:       mainId,
		ParentId: parentId,
	}

	return updatedCpp, nil
}
