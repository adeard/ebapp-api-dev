package popic

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindByPo(po string) ([]domain.PoPic, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindByPo(po string) ([]domain.PoPic, error) {
	var poPic []domain.PoPic

	q := r.db.Table("po_pic")

	if po != "" {
		q = q.Where("pekerjaan_no = ?", po)
	}

	err := q.Order("level asc").Find(&poPic).Error

	return poPic, err
}
