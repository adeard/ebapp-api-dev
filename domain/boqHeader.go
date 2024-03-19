package domain

import (
	"errors"
	"time"
)

type BoqHeader struct {
	RunNum            string    `json:"run_num" gorm:"column:run_num"`
	BoqNo             string    `json:"boq_no" gorm:"column:boq_no"`
	HeaderDescription string    `json:"header_description" gorm:"column:header_description"`
	HeaderVersion     string    `json:"header_version" gorm:"column:header_version"`
	HeaderStatus      bool      `json:"header_status" gorm:"column:header_status"`
	Created           time.Time `json:"created" gorm:"column:created"`
	CreatedBy         string    `json:"created_by" gorm:"column:created_by"`
	LastUpdated       time.Time `json:"last_updated" gorm:"column:lastupdated; default:nil"`
	LastUpdatedBy     string    `json:"last_updated_by" gorm:"column:lastupdatedby; default:null"`
	Category          string    `json:"category" gorm:"column:category"`
	Remarks           string    `json:"remarks" gorm:"column:remarks"`
}

type BoqHeaderRequest struct {
	BoqNo             string    `json:"boq_no"`
	HeaderDescription string    `json:"header_description"`
	HeaderVersion     string    `json:"header_version"`
	HeaderStatus      bool      `json:"header_status"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         string    `json:"created_by"`
	LastUpdated       time.Time `json:"last_updated"`
	LastUpdatedBy     string    `json:"last_updated_by"`
	Category          string    `json:"category"`
	Remarks           string    `json:"remarks"`
}

var (
	ErrNotFound = errors.New("not found")
)

type BoqClone struct {
	RunNum string `json:"run_num" gorm:"column:run_num"`
	BoqNo  string `json:"boq_no" gorm:"column:boq_no"`
}

type BoqHeaderResponse struct {
	Data    []BoqHeader `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}
