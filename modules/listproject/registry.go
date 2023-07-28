package listproject

import "gorm.io/gorm"

func ListProjectRegistry(db *gorm.DB) Service {
	listProjectRepository := NewRepository(db)
	listProjectService := NewService(listProjectRepository)

	return listProjectService
}
