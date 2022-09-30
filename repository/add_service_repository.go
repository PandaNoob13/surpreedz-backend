package repository

import (
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"

	"gorm.io/gorm"
)

type AddServiceRepository interface {
	AddService(newService *dto.ServiceDto) (int, error)
}

type addServiceRepository struct {
	db *gorm.DB
}

func (a *addServiceRepository) AddService(newService *dto.ServiceDto) (int, error) {
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
		return 0, err
	}

	toServicePrice := &model.ServicePrice{
		ServiceDetailId: toServiceDetail.ID,
		Price:           newService.Price,
	}
	if err := tx.Create(toServicePrice).Error; err != nil {
		tx.Rollback()
		return toServiceDetail.ID ,err
	}

	toVideoProfile := &model.VideoProfile{
		ServiceDetailId:  toServiceDetail.ID,
		VideoProfileLink: newService.VideoLink,
	}
	if err := tx.Create(toVideoProfile).Error; err != nil {
		tx.Rollback()
		return toServiceDetail.ID ,err
	}
	return toServiceDetail.ID ,tx.Commit().Error
}

func NewAddServiceRepository(db *gorm.DB) AddServiceRepository {
	repo := new(addServiceRepository)
	repo.db = db
	return repo
}
