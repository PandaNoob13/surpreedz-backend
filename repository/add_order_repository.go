package repository

import (
	"errors"
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"
	"time"

	"gorm.io/gorm"
)

type AddOrderRepository interface {
	AddOrder(newOrder *dto.OrderDto) error
}

type addOrderRepository struct {
	db *gorm.DB
}

func (a *addOrderRepository) AddOrder(newOrder *dto.OrderDto) error {
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	toOrder := &model.Order{
		BuyerId:         newOrder.BuyerId,
		ServiceDetailId: newOrder.ServiceDetailId,
		DueDate:         newOrder.DueDate,
	}
	if err := tx.Create(toOrder).Error; err != nil {
		tx.Rollback()
		return err
	}

	var order model.Order
	result := tx.Where("mst_order.buyer_id = ?", toOrder.BuyerId).Last(&order)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		} else {
			tx.Rollback()
			return err
		}
	}

	toOrderRequest := &model.OrderRequest{
		OrderId:       order.ID,
		Occasion:      newOrder.Occasion,
		RecipientName: newOrder.RecipientName,
		Message:       newOrder.Message,
		Description:   newOrder.RecipientDescription,
	}

	if err := tx.Create(toOrderRequest).Error; err != nil {
		tx.Rollback()
		return err
	}

	toOrderStatus := &model.OrderStatus{
		OrderId: order.ID,
		Status:  "Waiting for confirmation",
		Date:    time.Now(),
	}

	if err := tx.Create(toOrderStatus).Error; err != nil {
		tx.Rollback()
		return err
	}

	toPaymentStatus := &model.PaymentStatus{
		OrderId:       order.ID,
		StatusPayment: newOrder.StatusPayment,
	}

	if err := tx.Create(toPaymentStatus).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func NewAddOrderRepository(db *gorm.DB) AddOrderRepository {
	repo := new(addOrderRepository)
	repo.db = db
	return repo
}
