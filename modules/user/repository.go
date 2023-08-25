package user

import (
	"ebapp-api-dev/domain"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindByUserId(userId string) (domain.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUserId(userId string) (domain.User, error) {
	var user domain.User
	err := r.db.Table("user").Where("user_id = ?", userId).First(&user).Error
	fmt.Println(user)
	return user, err
}
