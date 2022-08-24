package usecase

import (
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
)

type UpdateServiceUseCase interface {
	EditService(serviceId int, existService *dto.EditServiceDto) error
}

type updateServiceUseCase struct {
	editServiceRepo repository.EditServiceRepository
}

func (u *updateServiceUseCase) EditService(serviceId int, existService *dto.EditServiceDto) error {
	return u.editServiceRepo.EditService(serviceId, existService)
}

func NewUpdateServiceUseCase(editServiceRepo repository.EditServiceRepository) UpdateServiceUseCase {
	return &updateServiceUseCase{
		editServiceRepo: editServiceRepo,
	}
}
