package domain

import "time"

type BoqHeader struct {
	RunNum            string    `json:"run_num"`
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

// type BoqHeaderResponse struct {
// 	RunNum            string `json:"run_num"`
// 	BoqNo             string `json:"boq_no"`
// 	HeaderDescription string `json:"header_description"`
// 	HeaderStatus      bool   `json:"header_status"`
// 	CreatedAt         string `json:"created_at"`
// 	CreatedBy         string `json:"created_by"`
// 	LastUpdated       string `json:"last_updated"`
// 	LastUpdatedBy     string `json:"last_updated_by"`
// 	Category          string `json:"category"`
// 	Remarks           string `json:"remarks"`
// }

type BoqHeaderResponse struct {
	Data    []BoqHeader `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}
