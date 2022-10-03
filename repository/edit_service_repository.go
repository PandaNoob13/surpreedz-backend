package repository

import (
	"errors"
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"

	"gorm.io/gorm"
)

type EditServiceRepository interface {
	EditService(serviceId int, existService *dto.EditServiceDto) error
}

type editServiceRepository struct {
	db *gorm.DB
}

func (e *editServiceRepository) EditService(serviceId int, existService *dto.EditServiceDto) error {
	tx := e.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var serviceDetail model.ServiceDetail
	result := tx.Where("mst_service_detail.id = ?", serviceId).Last(&serviceDetail)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		} else {
			tx.Rollback()
			return err
		}
	}

	toServiceDetail := model.ServiceDetail{
		ID: serviceId,
	}
	if err := tx.Model(&toServiceDetail).Updates(map[string]interface{}{
		"role":        existService.Role,
		"description": existService.Description,
	}).Error; err != nil {
		return err
	}

	toServicePrice := model.ServicePrice{
		ServiceDetailId: serviceId,
		Price:           existService.Price,
	}
	if err := tx.Create(&toServicePrice).Error; err != nil {
		tx.Rollback()
		return err
	}

	toVideoProfile := model.VideoProfile{
		ServiceDetailId:  serviceId,
		VideoProfileLink: existService.VideoLink,
	}
	if err := tx.Create(&toVideoProfile).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func NewEditServiceRepository(db *gorm.DB) EditServiceRepository {
	repo := new(editServiceRepository)
	repo.db = db
	return repo
}
