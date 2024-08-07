package poboqheader

import (
	"ebapp-api-dev/modules/poboqbody"

	"gorm.io/gorm"
)

func PoBoqHeaderRegistry(db *gorm.DB) Service {
	poBoqHeaderRepository := NewRepository(db)
	poBoqBodyService := poboqbody.PoBoqBodyRegistry(db)
	poBoqHeaderService := NewService(poBoqHeaderRepository, poBoqBodyService)

	return poBoqHeaderService
}
