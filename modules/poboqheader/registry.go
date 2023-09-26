package poboqheader

import "gorm.io/gorm"

func PoBoqHeaderRegistry(db *gorm.DB) Service {
	poBoqHeaderRepository := NewRepository(db)
	poBoqHeaderService := NewService(poBoqHeaderRepository)

	return poBoqHeaderService
}
