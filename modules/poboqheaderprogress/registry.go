package poboqheaderprogress

import (
	"ebapp-api-dev/modules/poboqbodyprogress"

	"gorm.io/gorm"
)

func PoBoqHeaderProgressRegistry(db *gorm.DB) Service {
	poBoqHeaderProgressRepository := NewRepository(db)
	poBoqBodyProgressService := poboqbodyprogress.PoBoqBodyProgressRegistry(db)
	poBoqHeaderProgressService := NewService(poBoqHeaderProgressRepository, poBoqBodyProgressService)

	return poBoqHeaderProgressService
}
