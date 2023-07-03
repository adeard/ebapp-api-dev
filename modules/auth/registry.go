package auth

import "gorm.io/gorm"

func AuthRegistry(db *gorm.DB) Service {
	authRepository := NewRepository(db)
	authService := NewService(authRepository)

	return authService
}
