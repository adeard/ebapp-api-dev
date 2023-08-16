package podatasapheader

import "gorm.io/gorm"

func PoDataSapHeaderRegistry(db *gorm.DB) Service {
	poDataSapHeaderRepository := NewRepository(db)
	poDataSapHeaderService := NewService(poDataSapHeaderRepository)

	return poDataSapHeaderService
}
