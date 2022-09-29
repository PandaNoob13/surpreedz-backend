package repository

import (
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"
	"time"

	"gorm.io/gorm"
)

type AddOrderStatusRepository interface {
	AddOrderStatus(newOrderStatus *dto.OrderStatusDto) error
}

type addOrderStatusRepository struct {
	db *gorm.DB
}

func (os *addOrderStatusRepository) AddOrderStatus(newOrderStatus *dto.OrderStatusDto) error {
	tx := os.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	toOrderStatus := &model.OrderStatus{
		OrderId: newOrderStatus.OrderId,
		Status:  newOrderStatus.Status,
		Date:    time.Now(),
	}
	if err := tx.Create(toOrderStatus).Error; err != nil {
		tx.Rollback()
		return err
	}

	// toRefund := &model.Refund{
	// 	OrderStatusId: toOrderStatus.ID,
	// 	Reason:        newOrderStatus.ResonOfRefund,
	// }
	// if err := tx.Create(toRefund).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	return tx.Commit().Error
}

func NewAddOrderStatusRepository(db *gorm.DB) AddOrderStatusRepository {
	repo := new(addOrderStatusRepository)
	repo.db = db
	return repo
}
