package poprogressheader

import "gorm.io/gorm"

func PoProgressHeaderRegistry(db *gorm.DB) Service {
	poProgressHeaderRepository := NewRepository(db)
	poProgressHeaderService := NewService(poProgressHeaderRepository)

	return poProgressHeaderService
}
