package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindOrderByIdUseCase interface {
	FindOrderById(id int) (model.Order, error)
}

type findOrderByIdUseCase struct {
	orderRepo repository.OrderRepository
}

func (f *findOrderByIdUseCase) FindOrderById(id int) (model.Order, error) {
	return f.orderRepo.FindById(id)
}

func NewFindOrderByIdUseCase(orderRepo repository.OrderRepository) FindOrderByIdUseCase {
	return &findOrderByIdUseCase{
		orderRepo: orderRepo,
	}
}
