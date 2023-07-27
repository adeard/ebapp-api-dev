package boqheader

import "gorm.io/gorm"

func BoqHeaderRegistry(db *gorm.DB) Service {
	boqHeaderRepository := NewRepository(db)
	boqHeaderService := NewService(boqHeaderRepository)

	return boqHeaderService
}
