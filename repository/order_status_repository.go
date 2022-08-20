package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type OrderStatusRepository interface {
	Create(orderStatus *model.OrderStatus) error
	FindById(id int) (model.OrderStatus, error)
	FindAll() ([]model.OrderStatus, error)
	UpdateByID(orderStatus *model.OrderStatus, by map[string]interface{}) error
	Delete(orderStatus *model.OrderStatus) error
}

type orderStatusRepository struct {
	db *gorm.DB
}

func (o *orderStatusRepository) Create(orderStatus *model.OrderStatus) error {
	result := o.db.Create(orderStatus).Error
	return result
}

func (o *orderStatusRepository) FindById(id int) (model.OrderStatus, error) {
	var orderStatus model.OrderStatus
	result := o.db.Preload("Refund").Where("order_id = ?", id).First(&orderStatus)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orderStatus, nil
		} else {
			return orderStatus, err
		}
	}
	return orderStatus, nil
}

func (o *orderStatusRepository) FindAll() ([]model.OrderStatus, error) {
	var orderStatus []model.OrderStatus
	result := o.db.Preload("Refund").Find(&orderStatus)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orderStatus, nil
		} else {
			return orderStatus, err
		}
	}
	return orderStatus, nil
}

func (o *orderStatusRepository) UpdateByID(orderStatus *model.OrderStatus, by map[string]interface{}) error {
	result := o.db.Model(orderStatus).Updates(by).Error
	return result
}

func (o *orderStatusRepository) Delete(orderStatus *model.OrderStatus) error {
	result := o.db.Delete(orderStatus).Error
	return result
}

func NewOrderStatusRepository(db *gorm.DB) OrderStatusRepository {
	repo := new(orderStatusRepository)
	repo.db = db
	return repo
}
