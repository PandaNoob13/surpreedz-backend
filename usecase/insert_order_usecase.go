package usecase

import (
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
)

type InsertOrderUseCase interface {
	AddOrder(newOrder dto.OrderDto) (string, error)
}

type insertOrderUseCase struct {
	addOrderRepo repository.AddOrderRepository
}

func (o *insertOrderUseCase) AddOrder(newOrder dto.OrderDto) (string, error) {
	return o.addOrderRepo.AddOrder(&newOrder)
}

func NewInsertOrderUseCase(addOrderRepo repository.AddOrderRepository) InsertOrderUseCase {
	return &insertOrderUseCase{
		addOrderRepo: addOrderRepo,
	}
}
