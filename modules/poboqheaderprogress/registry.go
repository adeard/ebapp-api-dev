package poboqheaderprogress

import "gorm.io/gorm"

func PoBoqHeaderProgressRegistry(db *gorm.DB) Service {
	poBoqHeaderProgressRepository := NewRepository(db)
	poBoqHeaderProgressService := NewService(poBoqHeaderProgressRepository)

	return poBoqHeaderProgressService
}
