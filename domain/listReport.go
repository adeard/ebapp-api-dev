package domain

import "time"

type ListReport struct {
	SpkNo           string     `json:"spk_no"`
	PoDate          time.Time  `json:"po_date"`
	Po              string     `json:"po"`
	PekerjaanName   string     `json:"pekerjaan_name"`
	UnitUsaha       string     `json:"unit_usaha"`
	Vendor          string     `json:"vendor"`
	Status          string     `json:"status"`
	Duration        string     `json:"duration"`
	StartDate       *time.Time `json:"start_date"`
	StartDateActual *time.Time `json:"start_date_actual" gorm:"column:start_date_actual"`
	EndDate         *time.Time `json:"end_date"`
	EndDateActual   *time.Time `json:"end_date_actual" gorm:"column:end_date_actual"`
	ProgActual      string     `json:"prog_actual"`
}

type ListReportResponse struct {
	Data    []ListReport `json:"data"`
	Status  int          `json:"status"`
	Message string       `json:"message"`
}
