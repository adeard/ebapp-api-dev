package user

import "gorm.io/gorm"

func UserRegistry(db *gorm.DB) Service {
	userRepository := NewRepository(db)
	userService := NewService(userRepository)

	return userService
}
