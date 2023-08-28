package user

import (
	"ebapp-api-dev/domain"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.UserRequest) ([]domain.User, error)
	FindByUserId(userId string) (domain.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.UserRequest) ([]domain.User, error) {
	var User []domain.User
	err := r.db.Table("user").Find(&User).Error
	return User, err
}

func (r *repository) FindByUserId(userId string) (domain.User, error) {
	var user domain.User
	err := r.db.Table("user").Where("user_id = ?", userId).First(&user).Error
	fmt.Println(user)
	return user, err
}
