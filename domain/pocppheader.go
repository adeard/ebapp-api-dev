package domain

import "time"

type PoCppHeader struct {
	RunNum         string    `json:"run_num" gorm:"column:run_num"`
	RunNumProgress string    `json:"run_num_progress" gorm:"column:run_num_progress"`
	Jenis          string    `json:"jenis" gorm:"column:jenis"`
	Date           time.Time `json:"date" gorm:"column:date;default:NULL"`
	StatusProgress string    `json:"status_progress" gorm:"column:status_progress"`
	LastUpdated    time.Time `json:"last_updated" gorm:"column:last_updated;default:NULL"`
}

type PoCppHeaderUpdate struct {
	Jenis          string    `json:"jenis" gorm:"column:jenis"`
	Date           time.Time `json:"date" gorm:"column:date;default:NULL"`
	StatusProgress string    `json:"status_progress" gorm:"column:status_progress"`
	LastUpdated    time.Time `json:"last_updated" gorm:"column:last_updated;default:NULL"`
}

type PoCppHeaderUpdateEbapp struct {
	Jenis       string    `json:"jenis" gorm:"column:jenis"`
	LastUpdated time.Time `json:"last_updated" gorm:"column:last_updated;"`
}

type PoCppHeaderUpdateStatus struct {
	StatusProgress string    `json:"status_progress" gorm:"column:status_progress"`
	LastUpdated    time.Time `json:"last_updated" gorm:"column:last_updated;"`
}

type PoCppHeaderResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []PoCppHeader `json:"data"`
}

type PoCppHeaderUpdateResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    []PoCppHeaderUpdate `json:"data"`
}

type PoCppHeaderUpdateEbappResponse struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    []PoCppHeaderUpdateEbapp `json:"data"`
}
