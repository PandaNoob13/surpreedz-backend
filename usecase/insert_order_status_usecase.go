package usecase

import (
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
)

type InsertOrderStatusUseCase interface {
	AddOrderStatus(newOrderStatus *dto.OrderStatusDto) error
}

type insertOrderStatusUseCase struct {
	addOrderStatusRepo repository.AddOrderStatusRepository
}

func (i *insertOrderStatusUseCase) AddOrderStatus(newOrderStatus *dto.OrderStatusDto) error {
	return i.addOrderStatusRepo.AddOrderStatus(newOrderStatus)
}

func NewInsertOrderStatusUseCase(addOrderStatusRepo repository.AddOrderStatusRepository) InsertOrderStatusUseCase {
	return &insertOrderStatusUseCase{
		addOrderStatusRepo: addOrderStatusRepo,
	}
}
