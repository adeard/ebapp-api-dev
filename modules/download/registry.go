package download

import "gorm.io/gorm"

func DownloadRegistry(db *gorm.DB) Service {
	downloadRepository := NewRepository(db)
	downloadService := NewService(downloadRepository)

	return downloadService
}
