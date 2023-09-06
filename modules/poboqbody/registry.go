package poboqbody

import "gorm.io/gorm"

func PoBoqBodyRegistry(db *gorm.DB) Service {
	poBoqBodyRepository := NewRepository(db)
	poBoqBodyService := NewService(poBoqBodyRepository)

	return poBoqBodyService
}
