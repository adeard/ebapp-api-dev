package domain

import "time"

type ListProject struct {
	Id              int       `json:"id"`
	Po              string    `json:"po"`
	PoDate          time.Time `json:"po_date"`
	PekerjaanNo     string    `json:"pekerjaan_no"`
	PekerjaanName   string    `json:"pekerjaan_name"`
	PekerjaanDate   time.Time `json:"pekerjaan_date"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	StartDateActual time.Time `json:"start_date_actual" gorm:"column:start_date_actual"`
	EndDateActual   time.Time `json:"end_date_actual" gorm:"column:end_date_actual"`
	UnitUsaha       string    `json:"unit_usaha"`
	Pt              string    `json:"pt"`
	Area            string    `json:"area"`
	Vendor          string    `json:"vendor"`
	Status          string    `json:"status"`
	Total           float64   `json:"total"`
	Currency        string    `json:"currency"`
}

type ListProjectRequest struct {
	Id              int       `json:"id"`
	Po              string    `json:"po"`
	PoDate          time.Time `json:"po_date"`
	PekerjaanNo     string    `json:"pekerjaan_no"`
	PekerjaanName   string    `json:"pekerjaan_name"`
	PekerjaanDate   time.Time `json:"pekerjaan_date"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	StartDateActual time.Time `json:"start_date_actual" gorm:"column:start_date_actual"`
	EndDateActual   time.Time `json:"end_date_actual" gorm:"column:end_date_actual"`
	UnitUsaha       string    `json:"unit_usaha"`
	Pt              string    `json:"pt"`
	Area            string    `json:"area"`
	Vendor          string    `json:"vendor"`
	Status          string    `json:"status"`
	Total           float64   `json:"total"`
	Currency        string    `json:"currency"`
}

type ListProjectsResponse struct {
	Data    []ListProject `json:"data"`
	Status  int           `json:"status"`
	Message string        `json:"message"`
}
