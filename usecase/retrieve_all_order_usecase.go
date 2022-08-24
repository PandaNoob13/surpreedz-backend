package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type RetrieveAllOrderUseCase interface {
	RetrieveAllOrder(page int, itemPerPage int) ([]model.Order, error)
}

type retrieveAllOrderUseCase struct {
	orderRepo repository.OrderRepository
}

func (r *retrieveAllOrderUseCase) RetrieveAllOrder(page int, itemPerPage int) ([]model.Order, error) {
	return r.orderRepo.FindAll(page, itemPerPage)
}

func NewRetrieveAllOrderUseCas(orderRepo repository.OrderRepository) RetrieveAllOrderUseCase {
	return &retrieveAllOrderUseCase{
		orderRepo: orderRepo,
	}
}
