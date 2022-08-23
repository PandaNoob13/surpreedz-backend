package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type UpdateServiceUseCase interface {
	EditService(existService *model.ServiceDetail, by map[string]interface{}) error
}

type updateServiceUseCase struct {
	serviceDetailRepo repository.ServiceDetailRepository
}

func (u *updateServiceUseCase) EditService(existService *model.ServiceDetail, by map[string]interface{}) error {
	return u.serviceDetailRepo.Update(existService, by)
}

func NewUpdateServiceUseCase(serviceDetailRepo repository.ServiceDetailRepository) UpdateServiceUseCase {
	return &updateServiceUseCase{
		serviceDetailRepo: serviceDetailRepo,
	}
}
