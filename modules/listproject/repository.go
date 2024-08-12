package listproject

import (
	"ebapp-api-dev/domain"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.ListProjectRequest) ([]domain.ListProject, error)
	FindById(id string) (domain.ListProject, error)
	FindByPlant(ids []string) ([]domain.ListProject, error)
	Store(input domain.ListProject) (domain.ListProject, error)
	Store2(input domain.ListProject2) (domain.ListProject2, error)
	Store3(input domain.ListProject3) (domain.ListProject3, error)
	Store4(input domain.ListProject4) (domain.ListProject4, error)

	FindByPekerjaanNo(id string) (domain.UpdateStatus, error)
	UpdateStatus(input domain.UpdateStatus) (domain.UpdateStatus, error)
	UpdateSpkNRetensi(input domain.UpdateSpkNRetensi, id string) (domain.UpdateSpkNRetensi, error)
	UpdatePlanningActualDate(input domain.ModelUpdateActualPlanningDate) error
	UpdateByPekerjaanNo(pekerjaanNo string, updateData map[string]interface{}) error
	SyncCanProgressFalse(pekerjaanNo string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.ListProjectRequest) ([]domain.ListProject, error) {
	var listProjects []domain.ListProject
	err := r.db.Table("list_project").Order("pekerjaan_date ASC").Find(&listProjects).Error
	return listProjects, err
}

func (r *repository) FindById(id string) (domain.ListProject, error) {
	var project domain.ListProject
	err := r.db.Table("list_project").Where("id =?", id).First(&project).Error
	return project, err
}

func (r *repository) FindByPlant(ids []string) ([]domain.ListProject, error) {
	var projects []domain.ListProject

	for _, id := range ids {
		var tmpProject []domain.ListProject
		parts := strings.Split(id, " ")

		err := r.db.Table("list_project").Where("unit_usaha LIKE ?", parts[0]+"%").Find(&tmpProject).Error
		for _, v := range tmpProject {
			projects = append(projects, v)
		}

		if err != nil {
			return nil, err
		}
	}
	return projects, nil
}

func (r *repository) Store(input domain.ListProject) (domain.ListProject, error) {
	err := r.db.Table("list_project").Create(&input).Error
	return input, err
}

func (r *repository) Store2(input domain.ListProject2) (domain.ListProject2, error) {
	err := r.db.Table("list_project").Create(&input).Error
	return input, err
}

func (r *repository) Store3(input domain.ListProject3) (domain.ListProject3, error) {
	err := r.db.Table("list_project").Create(&input).Error
	return input, err
}

func (r *repository) Store4(input domain.ListProject4) (domain.ListProject4, error) {
	err := r.db.Table("list_project").Create(&input).Error
	return input, err
}

func (r *repository) FindByPekerjaanNo(id string) (domain.UpdateStatus, error) {
	var project domain.UpdateStatus
	err := r.db.Table("list_project").Where("pekerjaan_no =?", id).First(&project).Error
	return project, err
}

func (r *repository) UpdateStatus(input domain.UpdateStatus) (domain.UpdateStatus, error) {
	err := r.db.Table("list_project").Where("pekerjaan_no =?", input.PekerjaanNo).Save(&input).Error
	return input, err
}

func (r *repository) UpdateSpkNRetensi(input domain.UpdateSpkNRetensi, id string) (domain.UpdateSpkNRetensi, error) {
	err := r.db.Table("list_project").Where("pekerjaan_no =?", input.PekerjaanNo).Where("[id] =?", id).Updates(map[string]interface{}{"spk_no": input.SpkNo, "retensi": input.Retensi, "masa_retensi": input.MasaRetensi}).Error
	return input, err
}

func (r *repository) UpdatePlanningActualDate(input domain.ModelUpdateActualPlanningDate) error {
	err := r.db.Table("list_project").Where("pekerjaan_no =?", input.PekerjaanNo).Updates(map[string]interface{}{"start_date": input.StartDate, "end_date": input.EndDate, "start_date_actual": input.StartDateActual, "end_date_actual": input.EndDateActual}).Error
	return err
}

func (r *repository) UpdateByPekerjaanNo(pekerjaanNo string, updateData map[string]interface{}) error {
	err := r.db.Debug().
		Table("list_project").
		Where("pekerjaan_no = ?", pekerjaanNo).
		Updates(updateData).Error

	return err
}

func (r *repository) SyncCanProgressFalse(pekerjaanNo string) error {
	err := r.db.Table("list_project").
		Where("pekerjaan_no = ?", pekerjaanNo).
		Update("can_progress", 0).Error
	return err
}
