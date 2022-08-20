package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type ServicePriceRepository interface {
	Insert(servicePrice *model.ServicePrice) error
	FindAllBy(preload string, condition string, searchValue ...interface{}) ([]model.ServicePrice, error)
	FindById(id int) (model.ServicePrice, error)
	RetrieveAll(page int, itemPerPage int) ([]model.ServicePrice, error)
	Update(servicePrice *model.ServicePrice, by map[string]interface{}) error
	Delete(id int) error
}

type servicePriceRepository struct {
	db *gorm.DB
}

func (sp *servicePriceRepository) Delete(id int) error {
	result := sp.db.Delete(&model.ServicePrice{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (sp *servicePriceRepository) Update(servicePrice *model.ServicePrice, by map[string]interface{}) error {
	result := sp.db.Model(servicePrice).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (sp *servicePriceRepository) RetrieveAll(page int, itemPerPage int) ([]model.ServicePrice, error) {
	var servicePrices []model.ServicePrice
	offset := itemPerPage * (page - 1)
	res := sp.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&servicePrices)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return servicePrices, nil
}

func (sp *servicePriceRepository) FindById(id int) (model.ServicePrice, error) {
	var servicePrice model.ServicePrice
	result := sp.db.First(&servicePrice, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return servicePrice, nil
		} else {
			return servicePrice, err
		}
	}
	return servicePrice, nil
}

func (sp *servicePriceRepository) FindAllBy(preload string, condition string, searcValue ...interface{}) ([]model.ServicePrice, error) {
	var servicePrices []model.ServicePrice
	if preload == "" {
		result := sp.db.Where(condition, searcValue...).Find(&servicePrices)
		if err := result.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			} else {
				return nil, err
			}
		}
	} else {
		result := sp.db.Preload(preload).Where(condition, searcValue...).Find(&servicePrices)
		if err := result.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			} else {
				return nil, err
			}
		}
	}
	return servicePrices, nil
}

func (sp *servicePriceRepository) Insert(servicePrice *model.ServicePrice) error {
	result := sp.db.Create(servicePrice)
	return result.Error
}

func NewServicePriceRepository(db *gorm.DB) ServicePriceRepository {
	repo := new(servicePriceRepository)
	repo.db = db
	return repo
}
