package poboqbodyprogress

import "gorm.io/gorm"

func PoBoqBodyProgressRegistry(db *gorm.DB) Service {
	poBoqBodyProgressRepository := NewRepository(db)
	poBoqBodyProgressService := NewService(poBoqBodyProgressRepository)

	return poBoqBodyProgressService
}
