package parentries

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.ParEntriesRequest) ([]domain.ParEntries, error)
	FindAllById(id string) ([]domain.ParEntries, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.ParEntriesRequest) ([]domain.ParEntries, error) {
	var parEntries []domain.ParEntries
	err := r.db.Table("par_entries").Find(&parEntries).Error
	return parEntries, err
}

func (r *repository) FindAllById(id string) ([]domain.ParEntries, error) {
	var parEntries []domain.ParEntries
	err := r.db.Table("par_entries").Where("tab_id =?", id).Find(&parEntries).Error
	return parEntries, err
}
