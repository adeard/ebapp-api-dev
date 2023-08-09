package domain

import "time"

type ListProject struct {
	Id              int       `json:"id"`
	PekerjaanNo     string    `json:"pekerjaan_no"`
	PekerjaanName   string    `json:"pekerjaan_name"`
	PekerjaanDate   time.Time `json:"pekerjaan_date"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	ActualStartDate time.Time `json:"start_date_actual"`
	ActualEndDate   time.Time `json:"end_date_actual"`
	UnitUsaha       string    `json:"unit_usaha"`
	Pt              string    `json:"pt"`
	Area            string    `json:"area"`
	Vendor          string    `json:"vendor"`
	Status          string    `json:"status"`
	Total           float64   `json:"total"`
}

type ListProjectRequest struct {
	PekerjaanNo   string    `json:"pekerjaan_no"`
	PekerjaanName string    `json:"pekerjaan_name"`
	Vendor        string    `json:"vendor"`
	Status        string    `json:"status"`
	PekerjaanDate time.Time `json:"pekerjaan_date"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Total         float64   `json:"total"`
}

type ListProjectsResponse struct {
	Data    []ListProject `json:"data"`
	Status  int           `json:"status"`
	Message string        `json:"message"`
}
