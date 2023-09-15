package poprojectaddendum

import "gorm.io/gorm"

func PoProjectAddendumRegistry(db *gorm.DB) Service {
	poProjectAddendumRepository := NewRepository(db)
	poProjectAddendumService := NewService(poProjectAddendumRepository)

	return poProjectAddendumService
}
