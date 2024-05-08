package poboqbodycpp

import "gorm.io/gorm"

func PoBoqBodyCppRegistry(db *gorm.DB) Service {
	poBoqBodyCppRepository := NewRepository(db)
	poBoqBodyCppService := NewService(poBoqBodyCppRepository)

	return poBoqBodyCppService
}
