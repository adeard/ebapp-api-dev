package report

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindByPlant(ids []string) ([]domain.ListReport, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByPlant(ids []string) ([]domain.ListReport, error) {
	var projects []domain.ListReport
	query := `SELECT a.spk_no, a.po_date, a.po, a.pekerjaan_name, a.unit_usaha, a.vendor, a.status, CAST(DATEDIFF(WEEK, a.start_date, a.end_date) AS VARCHAR) + ' weeks' as duration, a.start_date, a.start_date_actual, a.end_date, a.end_date_actual, 
CAST((b.new_prog + b.last_prog) as VARCHAR) + '%' as prog_actual
FROM list_project a
OUTER APPLY (
    SELECT TOP 1 *
    FROM po_progress_header b
    WHERE LEFT(b.run_num, LEN(b.run_num) - CHARINDEX('/', REVERSE(b.run_num))) = a.pekerjaan_no
    ORDER BY b.[date] DESC
) b where SUBSTRING (a.unit_usaha ,0 , CHARINDEX('-', a.unit_usaha)) in (?) ORDER BY a.start_date DESC
`
	err := r.db.Raw(query, ids).Find(&projects).Error
	return projects, err
}
