package poboqbody

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindByRunNum(runNum string, order string) ([]domain.PoBoqBody, error)
	Store(input domain.PoBoqBody) (domain.PoBoqBody, error)
	FindByItemNo(itemNo string) (domain.PoBoqBody, error)
	FindBoq(runNum string, order string, mainId string) ([]domain.PoBoqBody, error)
	Update(input domain.PoBoqBody) (domain.PoBoqBody, error)
	Delete(id string, orderId string, mainId string) error
	DeleteByOrder(id string, orderId string) error
	GetByRunNumAndOrder(runNum string, order string) ([]domain.PoBoqBody, error)
	GenerateMainId(runNum string, order string) (int, error)
	CheckItemNo(runNum string, order string, itemNo string) (bool, error)
	SelectMainId(runNum string, order string, itemNo string) (int, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByRunNum(runNum string, order string) ([]domain.PoBoqBody, error) {
	var boqBody []domain.PoBoqBody

	q := r.db.Table("po_boq_body")

	if runNum != "" {
		q = q.Where("run_num = ?", runNum).Where("[order] = ?", order)
	}

	err := q.Order("main_id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) Store(input domain.PoBoqBody) (domain.PoBoqBody, error) {
	err := r.db.Table("po_boq_body").Create(&input).Error
	return input, err
}

func (r *repository) FindByItemNo(itemNo string) (domain.PoBoqBody, error) {
	var poBoqBody domain.PoBoqBody
	err := r.db.Table("po_boq_body").Where("item_no = ?", itemNo).First(&poBoqBody).Error
	return poBoqBody, err
}

func (r *repository) FindByParentID(parentID int) ([]domain.PoBoqBody, error) {
	var boqBody []domain.PoBoqBody
	err := r.db.Table("po_boq_body").Where("parent_id = ?", parentID).Find(&boqBody).Error
	return boqBody, err
}

func (r *repository) FindBoq(runNum string, order string, mainId string) ([]domain.PoBoqBody, error) {
	var boqBody []domain.PoBoqBody

	q := r.db.Table("po_boq_body")

	if runNum != "" {
		q = q.Where("run_num = ?", runNum).Where("[order] = ?", order).Where("main_id = ?", mainId)
	}

	err := q.Order("main_id asc").Find(&boqBody).Error

	return boqBody, err
}

func (r *repository) Delete(id string, order string, mainId string) error {
	query := `
	WITH temp AS 
	(
		SELECT main_id, parent_id, 0 AS lvl
		FROM po_boq_body 
		WHERE main_id = ? AND run_num = ? AND [order] = ?

		UNION ALL 

		SELECT p.main_id, p.parent_id, C.lvl + 1 
		FROM temp C
		JOIN po_boq_body p ON C.main_id = p.parent_id 
	) 

	DELETE FROM po_boq_body
	WHERE main_id IN (SELECT main_id FROM temp);
	`

	err := r.db.Exec(query, mainId, id, order).Error

	return err
}

func (r *repository) DeleteByOrder(id string, orderId string) error {
	err := r.db.Table("po_boq_body").Where("run_num =?", id).Where("[order] =?", orderId).Delete(&domain.PoBoqBody{}).Error
	return err
}

func (r *repository) CheckItemNo(runNum string, order string, itemNo string) (bool, error) {
	var count int64

	err := r.db.Table("po_boq_body").
		Where("run_num = ?", runNum).
		Where("[order] = ?", order).
		Where("item_no = ?", itemNo).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func (r *repository) SelectMainId(runNum string, order string, itemNo string) (int, error) {
	var mainId int

	err := r.db.Table("po_boq_body").
		Select("main_id").
		Where("run_num = ?", runNum).
		Where("[order] = ?", order).
		Where("item_no = ?", itemNo).
		Order("main_id ASC").
		Limit(1).
		Scan(&mainId).Error

	if err != nil {
		return 0, err
	}

	return mainId, nil
}

func (r *repository) Update(input domain.PoBoqBody) (domain.PoBoqBody, error) {
	err := r.db.Table("po_boq_body").Where("run_num = ?", input.RunNum).Where("[order] = ?", input.Order).Where("main_id = ?", input.Id).Save(&input).Error

	if err != nil {
		return input, err
	}

	// Mengecek dan mengatasi kolom date yang kosong
	if input.StartDate == "" || input.EndDate == "" || input.StartDateActual == "" || input.EndDateActual == "" {
		updateCols := make(map[string]interface{})

		if input.StartDate == "" {
			updateCols["start_date"] = nil
		}

		if input.EndDate == "" {
			updateCols["end_date"] = nil
		}

		if input.StartDateActual == "" {
			updateCols["start_date_actual"] = nil
		}

		if input.EndDateActual == "" {
			updateCols["end_date_actual"] = nil
		}

		err := r.db.Table("po_boq_body").
			Where("run_num = ?", input.RunNum).
			Where("[order] = ?", input.Order).
			Where("main_id = ?", input.Id).
			Updates(updateCols).Error

		if err != nil {
			return input, err
		}
	}

	return input, err
}

func (r *repository) GetByRunNumAndOrder(runNum string, order string) ([]domain.PoBoqBody, error) {
	result := []domain.PoBoqBody{}

	err := r.db.Table("po_boq_body").Where("run_num =?", runNum).Where("[order] =?", order).Find(&result).Error
	return result, err
}

func (r *repository) GenerateMainId(runNum string, order string) (int, error) {
	var maxValue int

	err := r.db.Raw(`
		SELECT TOP 1 
			CASE 
				WHEN CAST(main_id AS int) > parent_id THEN CAST(main_id AS int)
				ELSE parent_id
			END + 1 AS max_value
		FROM DB_eBAPP.dbo.po_boq_body
		WHERE run_num = ? 
		AND [order] = ?
		ORDER BY max_value DESC
	`, runNum, order).Scan(&maxValue).Error

	return maxValue, err
}
