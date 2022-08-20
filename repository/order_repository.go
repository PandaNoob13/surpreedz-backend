package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *model.Order) error
	FindById(id int) (model.Order, error)
	FindAll() ([]model.Order, error)
	UpdateByID(order *model.Order, by map[string]interface{}) error
	Delete(order *model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func (o *orderRepository) Create(order *model.Order) error {
	result := o.db.Create(order).Error
	return result
}

func (o *orderRepository) FindById(id int) (model.Order, error) {
	var order model.Order
	result := o.db.Preload("OrderStatus.Refund").Preload("OrderRequest").Preload("Feedback").Preload("VideoResult").Where("mst_order.id = ?", id).First(&order)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, nil
		} else {
			return order, err
		}
	}
	return order, nil
}

func (o *orderRepository) FindAll() ([]model.Order, error) {
	var order []model.Order
	result := o.db.Preload("OrderStatus.Refund").Preload("OrderRequest").Preload("Feedback").Preload("VideoResult").Find(&order)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, nil
		} else {
			return order, err
		}
	}
	return order, nil
}

func (o *orderRepository) UpdateByID(order *model.Order, by map[string]interface{}) error {
	result := o.db.Model(order).Updates(by).Error
	return result
}

func (o *orderRepository) Delete(order *model.Order) error {
	result := o.db.Delete(order).Error
	return result
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	repo := new(orderRepository)
	repo.db = db
	return repo
}
