package usecase

import (
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
)

type ShowServicesHomePageUseCase interface {
	HomePageRetrieveAll(page int, itemPerPage int) ([]dto.AccountCreateDto, error)
}

type showServicesHomePageUseCase struct {
	serviceDetailRepo repository.ServiceDetailRepository
}

func (h *showServicesHomePageUseCase) HomePageRetrieveAll(page int, itemPerPage int) ([]dto.AccountCreateDto, error) {
	return h.serviceDetailRepo.HomePageRetrieveAll(page, itemPerPage)
}

func NewShowServiceHomePageUseCase(serviceDetailRepo repository.ServiceDetailRepository) ShowServicesHomePageUseCase {
	return &showServicesHomePageUseCase{
		serviceDetailRepo: serviceDetailRepo,
	}
}
