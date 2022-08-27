package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *model.Order) error
	FindById(id int) (model.Order, error)
	FindAll(page int, itemPerPage int) ([]model.Order, error)
	FindByBuyerId(id int) (model.Order, error)
	FindAllByBuyerId(buyerId int) ([]model.Order, error)
	FindAllByServiceDetailId(serviceDetailId int) ([]model.Order, error)
	UpdateByID(order *model.Order, by map[string]interface{}) error
	Delete(order *model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func (o *orderRepository) FindAllByBuyerId(buyerId int) ([]model.Order, error) {
	var order []model.Order
	result := o.db.Where("mst_order.buyer_id = ?", buyerId).Find(&order)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, nil
		} else {
			return order, err
		}
	}
	return order, nil
}

func (o *orderRepository) FindAllByServiceDetailId(serviceDetailId int) ([]model.Order, error) {
	var order []model.Order
	result := o.db.Where("mst_order.service_detail_id = ?", serviceDetailId).Find(&order)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, nil
		} else {
			return order, err
		}
	}
	return order, nil
}

func (o *orderRepository) FindByBuyerId(id int) (model.Order, error) {
	var order model.Order
	result := o.db.Where("mst_order.buyer_id = ?", id).Last(&order)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, nil
		} else {
			return order, err
		}
	}
	return order, nil
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

func (o *orderRepository) FindAll(page int, itemPerPage int) ([]model.Order, error) {
	var order []model.Order
	offset := itemPerPage * (page - 1)
	result := o.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Preload("OrderStatus.Refund").Preload("OrderRequest").Preload("Feedback").Preload("VideoResult").Find(&order)
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
