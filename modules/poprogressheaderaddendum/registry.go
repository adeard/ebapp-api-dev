package poprogressheaderaddendum

import "gorm.io/gorm"

func PoProgressHeaderAddendumRegistry(db *gorm.DB) Service {
	poProgressHeaderAddendumRepository := NewRepository(db)
	poProgressHeaderAddendumService := NewService(poProgressHeaderAddendumRepository)

	return poProgressHeaderAddendumService
}
