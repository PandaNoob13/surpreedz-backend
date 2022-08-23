package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindServiceUseCase interface {
	FindServiceById(id int) (model.ServiceDetail, error)
}

type findServiceUseCase struct {
	serviceDetailRepo repository.ServiceDetailRepository
}

func (f *findServiceUseCase) FindServiceById(id int) (model.ServiceDetail, error) {
	return f.serviceDetailRepo.FindById(id)
}

func NewFindServiceUseCase(serviceDetailRepo repository.ServiceDetailRepository) FindServiceUseCase {
	return &findServiceUseCase{
		serviceDetailRepo: serviceDetailRepo,
	}
}
