package domain

import "time"

type ListProject struct {
	Id            int       `json:"id"`
	PekerjaanNo   string    `json:"pekerjaan_no"`
	PekerjaanName string    `json:"pekerjaan_name"`
	Vendor        string    `json:"vendor"`
	Status        string    `json:"status"`
	PekerjaanDate time.Time `json:"pekerjaan_date"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Total         float64   `json:"total"`
}
