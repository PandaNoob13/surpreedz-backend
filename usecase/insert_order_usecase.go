package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
	"surpreedz-backend/utils"
)

type InsertOrderUseCase interface {
	AddOrder(buyerId int, serviceId int, dueDate string, occasion string, recipientName string, message string, description string) error
}

type insertOrderUseCase struct {
	orderRepo        repository.OrderRepository
	orderRequestRepo repository.OrderRequestRepository
}

func (o *insertOrderUseCase) AddOrder(buyerId int, serviceId int, dueDate string, occasion string, recipientName string, message string, description string) error {
	insertOrder := dto.OrderDto{
		BuyerId:              buyerId,
		ServiceDetailId:      serviceId,
		DueDate:              dueDate,
		Occasion:             occasion,
		RecipientName:        recipientName,
		Message:              message,
		RecipientDescription: description,
	}

	toOrder := model.Order{
		BuyerId:         buyerId,
		ServiceDetailId: serviceId,
		DueDate:         insertOrder.DueDate,
	}
	err := o.orderRepo.Create(&toOrder)
	utils.IsError(err)

	order, err := o.orderRepo.FindByBuyerId(buyerId)
	utils.IsError(err)

	toOrderRequest := model.OrderRequest{
		OrderId:       order.ID,
		Occasion:      insertOrder.Occasion,
		RecipientName: insertOrder.RecipientName,
		Message:       insertOrder.Message,
		Description:   insertOrder.RecipientDescription,
	}
	err1 := o.orderRequestRepo.Create(&toOrderRequest)
	utils.IsError(err1)

	return nil
}

func NewInsertOrderUseCase(orderRepo repository.OrderRepository, orderRequestRepo repository.OrderRequestRepository) InsertOrderUseCase {
	return &insertOrderUseCase{
		orderRepo:        orderRepo,
		orderRequestRepo: orderRequestRepo,
	}
}
