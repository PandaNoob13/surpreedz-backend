package usecase

import (
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
)

type InsertServiceUseCase interface {
	AddService(newService *dto.ServiceDto) error
}

type insertServiceUseCase struct {
	addServiceRepo repository.AddServiceRepository
}

func (s *insertServiceUseCase) AddService(newService *dto.ServiceDto) error {
	return s.addServiceRepo.AddService(newService)
}

func NewInsertServiceUseCase(addServiceRepo repository.AddServiceRepository) InsertServiceUseCase {
	return &insertServiceUseCase{
		addServiceRepo: addServiceRepo,
	}
}
