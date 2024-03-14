package progressattachment

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	GetAttachment(id string) ([]domain.ProgressAttachment, error)
	Store(input domain.ProgressAttachment) (domain.ProgressAttachment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAttachment(id string) ([]domain.ProgressAttachment, error) {
	var progress_attachment []domain.ProgressAttachment

	q := r.db.Table("progress_attachment")

	if id != "" {
		q = q.Where("run_num = ?", id)
	}

	err := q.Order("date asc").Find(&progress_attachment).Error

	return progress_attachment, err
}

func (r *repository) Store(input domain.ProgressAttachment) (domain.ProgressAttachment, error) {
	err := r.db.Table("progress_attachment").Create(&input).Error
	return input, err
}
