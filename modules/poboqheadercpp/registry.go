package poboqheadercpp

import "gorm.io/gorm"

func PoBoqHeaderCppRegistry(db *gorm.DB) Service {
	poBoqHeaderCppRepository := NewRepository(db)
	poBoqHeaderCppService := NewService(poBoqHeaderCppRepository)

	return poBoqHeaderCppService
}
