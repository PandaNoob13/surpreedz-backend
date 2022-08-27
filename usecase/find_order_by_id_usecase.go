package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindOrderByIdUseCase interface {
	FindOrderById(id int) (model.Order, error)
	FindAllOrderByBuyerId(id int) ([]model.Order, error)
	FindAllOrderByServiceDetailId(id int) ([]model.Account, error)
}

type findOrderByIdUseCase struct {
	orderRepo repository.OrderRepository
}

func (f *findOrderByIdUseCase) FindOrderById(id int) (model.Order, error) {
	return f.orderRepo.FindById(id)
}

func (f *findOrderByIdUseCase) FindAllOrderByBuyerId(id int) ([]model.Order, error) {
	return f.orderRepo.FindAllByBuyerId(id)
}

func (f *findOrderByIdUseCase) FindAllOrderByServiceDetailId(id int) ([]model.Account, error) {
	return f.orderRepo.FindAllByServiceDetailId(id)
}

func NewFindOrderByIdUseCase(orderRepo repository.OrderRepository) FindOrderByIdUseCase {
	return &findOrderByIdUseCase{
		orderRepo: orderRepo,
	}
}
