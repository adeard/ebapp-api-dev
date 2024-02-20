package domain

import "time"

type PoProgressHeaderAddendum struct {
	RunNum   string    `json:"run_num" gorm:"column:run_num"`
	Date     time.Time `json:"date" gorm:"column:date;default:NULL"`
	Po       string    `json:"status" gorm:"column:po"`
	LastProg float32   `json:"last_prog" gorm:"last_prog"`
	NewProg  float32   `json:"new_prog" gorm:"new_prog"`
}

type PoProgressHeaderAddendumResponse struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    []PoProgressHeaderAddendum `json:"data"`
}
