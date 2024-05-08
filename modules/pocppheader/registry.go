package pocppheader

import "gorm.io/gorm"

func PoCppHeaderRegistry(db *gorm.DB) Service {
	poCppHeaderRepository := NewRepository(db)
	poCppHeaderService := NewService(poCppHeaderRepository)

	return poCppHeaderService
}
