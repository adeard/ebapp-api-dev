package poproject

import "gorm.io/gorm"

func PoProjectRegistry(db *gorm.DB) Service {
	poProjectRepository := NewRepository(db)
	poProjectService := NewService(poProjectRepository)

	return poProjectService
}
