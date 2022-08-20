package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type OrderRequestRepository interface {
	Create(orderRequest *model.OrderRequest) error
	FindById(id int) (model.OrderRequest, error)
	FindAll() ([]model.OrderRequest, error)
	UpdateByID(orderRequest *model.OrderRequest, by map[string]interface{}) error
	Delete(orderRequest *model.OrderRequest) error
}

type orderRequestRepository struct {
	db *gorm.DB
}

func (o *orderRequestRepository) Create(orderRequest *model.OrderRequest) error {
	result := o.db.Create(orderRequest).Error
	return result
}

func (o *orderRequestRepository) FindById(id int) (model.OrderRequest, error) {
	var orderRequest model.OrderRequest
	result := o.db.First(&orderRequest, "order_id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orderRequest, nil
		} else {
			return orderRequest, err
		}
	}
	return orderRequest, nil
}

func (o *orderRequestRepository) FindAll() ([]model.OrderRequest, error) {
	var orderRequests []model.OrderRequest
	result := o.db.Find(&orderRequests)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orderRequests, nil
		} else {
			return orderRequests, err
		}
	}
	return orderRequests, nil
}

func (o *orderRequestRepository) UpdateByID(orderRequest *model.OrderRequest, by map[string]interface{}) error {
	result := o.db.Model(orderRequest).Updates(by).Error
	return result
}

func (o *orderRequestRepository) Delete(orderRequest *model.OrderRequest) error {
	result := o.db.Delete(orderRequest).Error
	return result
}

func NewOrderRequestRepository(db *gorm.DB) OrderRequestRepository {
	repo := new(orderRequestRepository)
	repo.db = db
	return repo
}
