package domain

import "time"

type PoProgressHeader struct {
	RunNum   string    `json:"run_num" gorm:"column:run_num"`
	Date     time.Time `json:"date" gorm:"column:date;default:NULL"`
	Status   string    `json:"status" gorm:"column:status"`
	IsEbapp  int       `json:"is_ebapp" gorm:"column:isebapp"`
	LastProg int       `json:"last_prog" gorm:"last_prog"`
	NewProg  int       `json:"new_prog" gorm:"new_prog"`
}

type PoProgressHeaderResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    []PoProgressHeader `json:"data"`
}
