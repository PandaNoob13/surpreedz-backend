package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindOrderCMSUseCase interface {
	FindOrderWithCondition(table string, condition string) ([]model.Order, error)
}

type findOrderCMSUseCase struct {
	cmsOrderRepo repository.CMSOrderRepository
}

func (f *findOrderCMSUseCase) FindOrderWithCondition(table string, condition string) ([]model.Order, error) {
	return f.cmsOrderRepo.GetAllOrder(table, condition)
}

func NewFindOrderCMSUseCase(cmsOrderRepo repository.CMSOrderRepository) FindOrderCMSUseCase {
	return &findOrderCMSUseCase{
		cmsOrderRepo: cmsOrderRepo,
	}
}
