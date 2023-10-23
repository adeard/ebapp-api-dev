package popic

import "gorm.io/gorm"

func PoPicRegistry(db *gorm.DB) Service {
	poPicRepository := NewRepository(db)
	poPicService := NewService(poPicRepository)

	return poPicService
}
