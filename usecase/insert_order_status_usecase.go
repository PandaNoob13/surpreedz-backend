package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
	"surpreedz-backend/utils"
)

type InsertOrderStatusUseCase interface {
	AddOrderStatus(orderStatusId int, orderId int, status string, reason string) error
}

type insertOrderStatusUseCase struct {
	orderStatusRepo repository.OrderStatusRepository
	refundRepo      repository.RefundRepository
}

func (i *insertOrderStatusUseCase) AddOrderStatus(orderStatusId int, orderId int, status string, reason string) error {
	insertOrderStatus := dto.OrderStatusDto{
		OrderId:       orderId,
		Status:        status,
		ResonOfRefund: reason,
	}

	toOrderStatus := model.OrderStatus{
		OrderId: insertOrderStatus.OrderId,
		Status:  insertOrderStatus.Status,
	}
	err := i.orderStatusRepo.Create(&toOrderStatus)
	utils.IsError(err)

	toRefund := model.Refund{
		OrderStatusId: orderStatusId,
		Reason:        insertOrderStatus.ResonOfRefund,
	}
	err1 := i.refundRepo.Create(&toRefund)
	utils.IsError(err1)

	return nil
}

func NewInsertOrderStatusUseCase(orderStatusRepo repository.OrderStatusRepository, refundRepo repository.RefundRepository) InsertOrderStatusUseCase {
	return &insertOrderStatusUseCase{
		orderStatusRepo: orderStatusRepo,
		refundRepo:      refundRepo,
	}
}
