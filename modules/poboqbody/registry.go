package poboqbody

import (
	"ebapp-api-dev/modules/listproject"

	"gorm.io/gorm"
)

func PoBoqBodyRegistry(db *gorm.DB) Service {
	poBoqBodyRepository := NewRepository(db)
	listProjectService := listproject.ListProjectRegistry(db)
	poBoqBodyService := NewService(poBoqBodyRepository, listProjectService)

	return poBoqBodyService
}
