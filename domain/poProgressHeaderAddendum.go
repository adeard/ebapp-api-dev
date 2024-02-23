package domain

import "time"

type PoProgressHeaderAddendum struct {
	RunNum      string    `json:"run_num" gorm:"column:run_num"`
	Date        time.Time `json:"date" gorm:"column:date;default:NULL"`
	Po          string    `json:"po" gorm:"column:po"`
	LastProg    float32   `json:"last_prog" gorm:"last_prog"`
	NewProg     float32   `json:"new_prog" gorm:"new_prog"`
	LastUpdated time.Time `json:"last_updated" gorm:"column:last_updated;default:NULL"`
}

type PoProgressHeaderAddendumUpdate struct {
	Po          string    `json:"po" gorm:"column:po"`
	LastProg    float32   `json:"last_prog" gorm:"column:last_prog"`
	NewProg     float32   `json:"new_prog" gorm:"column:new_prog"`
	LastUpdated time.Time `json:"last_updated" gorm:"column:last_updated;default:NULL"`
}

type PoProgressHeaderAddendumResponse struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    []PoProgressHeaderAddendum `json:"data"`
}

type PoProgressHeaderAddendumResponseUpdate struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    []PoProgressHeaderAddendumUpdate `json:"data"`
}
