package repository

import (
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"

	"gorm.io/gorm"
)

type AddServiceRepository interface {
	AddService(newService *dto.ServiceDto) error
}

type addServiceRepository struct {
	db *gorm.DB
}

func (a *addServiceRepository) AddService(newService *dto.ServiceDto) error {
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	toServiceDetail := &model.ServiceDetail{
		SellerId:    newService.SellerId,
		Role:        newService.Role,
		Description: newService.Description,
	}
	if err := tx.Create(toServiceDetail).Error; err != nil {
		tx.Rollback()
		return err
	}

	toServicePrice := &model.ServicePrice{
		ServiceDetailId: toServiceDetail.ID,
		Price:           newService.Price,
	}
	if err := tx.Create(toServicePrice).Error; err != nil {
		tx.Rollback()
		return err
	}

	toVideoProfile := &model.VideoProfile{
		ServiceDetailId:  toServiceDetail.ID,
		VideoProfileLink: newService.VideoLink,
	}
	if err := tx.Create(toVideoProfile).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func NewAddServiceRepository(db *gorm.DB) AddServiceRepository {
	repo := new(addServiceRepository)
	repo.db = db
	return repo
}
