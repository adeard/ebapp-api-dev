package poboqheadercpp

import (
	"ebapp-api-dev/modules/poboqbodycpp"

	"gorm.io/gorm"
)

func PoBoqHeaderCppRegistry(db *gorm.DB) Service {
	poBoqHeaderCppRepository := NewRepository(db)
	poBoqBodyCppService := poboqbodycpp.PoBoqBodyCppRegistry(db)
	poBoqHeaderCppService := NewService(poBoqHeaderCppRepository, poBoqBodyCppService)

	return poBoqHeaderCppService
}
