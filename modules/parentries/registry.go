package parentries

import "gorm.io/gorm"

func ParEntriesRegistry(db *gorm.DB) Service {
	parEntriesRepository := NewRepository(db)
	parEntriesService := NewService(parEntriesRepository)

	return parEntriesService
}
