package boqbody

import "gorm.io/gorm"

func BoqBodyRegistry(db *gorm.DB) Service {
	boqBodyRepository := NewRepository(db)
	boqBodyService := NewService(boqBodyRepository)

	return boqBodyService
}
