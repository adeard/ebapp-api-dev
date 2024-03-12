package progressattachment

import "gorm.io/gorm"

func ProgressAttachmentRegistry(db *gorm.DB) Service {
	progressAttachmentRepository := NewRepository(db)
	progressAttachmentService := NewService(progressAttachmentRepository)

	return progressAttachmentService
}
