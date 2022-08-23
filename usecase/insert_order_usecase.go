package usecase

import (
	"surpreedz-backend/dto"
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
	"surpreedz-backend/utils"
)

type InsertOrderUseCase interface {
	AddOrder(buyerId int, serviceId int, orderId int, dueDate string, occasion string, recipientName string, message string, description string) error
}

type insertOrderUseCase struct {
	orderRepo        repository.OrderRepository
	orderRequestRepo repository.OrderRequestRepository
}

func (o *insertOrderUseCase) AddOrder(orderId int, buyerId, serviceId int, dueDate string, occasion string, recipientName string, message string, description string) error {
	insertOrder := dto.OrderDto{
		BuyerId:              buyerId,
		ServiceDetailId:      serviceId,
		DueDate:              dueDate,
		Occassion:            occasion,
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

	toOrderRequest := model.OrderRequest{
		OrderId:       orderId,
		Ocassion:      insertOrder.Occassion,
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
