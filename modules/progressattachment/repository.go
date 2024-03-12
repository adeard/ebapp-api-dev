package progressattachment

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.ProgressAttachment) (domain.ProgressAttachment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Store(input domain.ProgressAttachment) (domain.ProgressAttachment, error) {
	err := r.db.Table("progress_attachment").Create(&input).Error
	return input, err
}
