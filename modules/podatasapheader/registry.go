package podatasapheader

import "gorm.io/gorm"

func PoDataSapHeaderRegistry(db2 *gorm.DB) Service {
	poDataSapHeaderRepository := NewRepository(db2)
	poDataSapHeaderService := NewService(poDataSapHeaderRepository)

	return poDataSapHeaderService
}
