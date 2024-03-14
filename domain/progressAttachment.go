package domain

import "time"

type ProgressAttachment struct {
	RunNum   string    `json:"run_num" gorm:"column:run_num"`
	FilePath string    `json:"file_path" gorm:"column:file_path"`
	Date     time.Time `json:"date" gorm:"column:date;default:NULL"`
	UploadBy string    `json:"upload_by" gorm:"upload_by"`
	FileName string    `json:"file_name" gorm:"column:file_name"`
}

type ProgressAttachmentView struct {
	Source   string    `json:"source"`
	FileName string    `json:"file_name" gorm:"column:file_name"`
	UploadBy string    `json:"upload_by" gorm:"upload_by"`
	Date     time.Time `json:"date" gorm:"column:date;default:NULL"`
}

type ProgressAttachmentResponse struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    []ProgressAttachment `json:"data"`
}

type ProgressAttachmentViewResponse struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    []ProgressAttachmentView `json:"data"`
}
