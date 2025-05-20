package report

import "gorm.io/gorm"

func ReportRegistry(db *gorm.DB) Service {
	ReportRepository := NewRepository(db)
	ReportService := NewService(ReportRepository)

	return ReportService
}
