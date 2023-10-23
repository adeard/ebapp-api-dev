package popic

import (
	"ebapp-api-dev/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindPic(uid string, po string, role string) (domain.PoPic, error)
	FindPicByLevel(uid string, po string, level int) (domain.PoPic, error)
	FindByPo(po string) ([]domain.PoPic, error)
	Store(input domain.PoPic) (domain.PoPic, error)
	Update(input domain.PoPic) (domain.PoPic, error)
	Delete(uid string, po string, level int) (domain.PoPic, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindPic(uid string, po string, role string) (domain.PoPic, error) {
	var pic domain.PoPic
	err := r.db.Table("po_pic").Where("uid =?", uid).Where("pekerjaan_no =?", po).Where("role =?", role).First(&pic).Error
	return pic, err
}

func (r *repository) FindPicByLevel(uid string, po string, level int) (domain.PoPic, error) {
	var pic domain.PoPic
	err := r.db.Table("po_pic").Where("uid=?", uid).Where("pekerjaan_no =?", po).Where("level =?", level).First(&pic).Error
	return pic, err
}

func (r *repository) FindByPo(po string) ([]domain.PoPic, error) {
	var poPic []domain.PoPic

	q := r.db.Table("po_pic")

	if po != "" {
		q = q.Where("pekerjaan_no = ?", po)
	}

	err := q.Order("level asc").Find(&poPic).Error

	return poPic, err
}

func (r *repository) Store(input domain.PoPic) (domain.PoPic, error) {
	err := r.db.Table("po_pic").Create(&input).Error
	return input, err
}

func (r *repository) Delete(uid string, po string, level int) (domain.PoPic, error) {
	var pic domain.PoPic
	err := r.db.Table("po_pic").Where("uid =?", uid).Where("pekerjaan_no =?", po).Where("level =?", level).Delete(&pic).Error
	return pic, err
}

func (r *repository) Update(input domain.PoPic) (domain.PoPic, error) {
	err := r.db.Table("po_pic").Where("pekerjaan_no =?", input.RunNum).Where("uid =?", input.Uid).Where("level =?", input.Id).Save(&input).Error
	return input, err
}
