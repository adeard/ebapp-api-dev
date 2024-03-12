package domain

import "time"

type progressAttachment struct {
	RunNum   string    `json:"run_num" gorm:"column:run_num"`
	FilePath string    `json:"file_path" gorm:"column:file_path"`
	Date     time.Time `json:"date" gorm:"column:date;default:NULL"`
	UploadBy string    `json:"upload_by" gorm:"upload_by"`
}

type progressAttachmentResponse struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    []progressAttachment `json:"data"`
}
