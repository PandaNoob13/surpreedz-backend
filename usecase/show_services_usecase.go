package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type ShowServicesHomePageUseCase interface {
	HomePageRetrieveAll(page int, itemPerPage int) ([]model.ServiceDetail, error)
}

type showServicesHomePageUseCase struct {
	serviceDetailRepo repository.ServiceDetailRepository
}

func (h *showServicesHomePageUseCase) HomePageRetrieveAll(page int, itemPerPage int) ([]model.ServiceDetail, error) {
	return h.serviceDetailRepo.HomePageRetrieveAll(page, itemPerPage)
}

func NewShowServiceHomePageUseCase(serviceDetailRepo repository.ServiceDetailRepository) ShowServicesHomePageUseCase {
	return &showServicesHomePageUseCase{
		serviceDetailRepo: serviceDetailRepo,
	}
}
