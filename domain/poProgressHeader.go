package domain

import "time"

type PoProgressHeader struct {
	RunNum      string    `json:"run_num" gorm:"column:run_num"`
	Date        time.Time `json:"date" gorm:"column:date;default:NULL"`
	Status      string    `json:"status" gorm:"column:status"`
	IsEbapp     int       `json:"is_ebapp" gorm:"column:isebapp"`
	LastProg    float32   `json:"last_prog" gorm:"last_prog"`
	NewProg     float32   `json:"new_prog" gorm:"new_prog"`
	Lock        int       `json:"lock" gorm:"lock"`
	LastUpdated time.Time `json:"last_updated" gorm:"column:last_updated;default:NULL"`
}

type PoProgressHeaderUpdate struct {
	Status      string    `json:"status" gorm:"column:status;default:NULL"`
	IsEbapp     int       `json:"is_ebapp" gorm:"column:isebapp;default:NULL"`
	LastProg    float32   `json:"last_prog" gorm:"last_prog;default:NULL"`
	NewProg     float32   `json:"new_prog" gorm:"new_prog;default:NULL"`
	LastUpdated time.Time `json:"last_updated" gorm:"column:last_updated;default:NULL"`
}

type PoProgressHeaderResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    []PoProgressHeader `json:"data"`
}

type PoProgressHeaderUpdateResponse struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    []PoProgressHeaderUpdate `json:"data"`
}
