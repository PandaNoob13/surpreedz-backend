package usecase

import (
	"surpreedz-backend/dto"
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
	"surpreedz-backend/utils"
)

type InsertServiceUseCase interface {
	AddService(serviceDetailId int, accountId int, role string, description string, price int, videoLink string) error
}

type insertServiceUseCase struct {
	serviceDetailRepo repository.ServiceDetailRepository
	servicePriceRepo  repository.ServicePriceRepository
	videoProfileRepo  repository.VideoProfileRepository
}

func (s *insertServiceUseCase) AddService(serviceDetailId int, accountId int, role string, description string, price int, videoLink string) error {
	insertService := dto.ServiceDto{
		SellerId:    accountId,
		Role:        role,
		Description: description,
		Price:       price,
		VideoLink:   videoLink,
	}

	toServiceDetail := model.ServiceDetail{
		SellerId:    insertService.SellerId,
		Role:        insertService.Role,
		Description: insertService.Description,
	}
	err1 := s.serviceDetailRepo.Insert(&toServiceDetail)
	utils.IsError(err1)

	toServicePrice := model.ServicePrice{
		ServiceDetailId: serviceDetailId,
		Price:           insertService.Price,
	}
	err2 := s.servicePriceRepo.Insert(&toServicePrice)
	utils.IsError(err2)

	toVideoProfile := model.VideoProfile{
		ServiceDetailId:  serviceDetailId,
		VideoProfileLink: insertService.VideoLink,
	}
	err3 := s.videoProfileRepo.Insert(&toVideoProfile)
	utils.IsError(err3)
	return nil
}

func NewInsertServiceUseCase(accountDetailRepo repository.AccountDetailRepository, serviceDetailRepo repository.ServiceDetailRepository, servicePriceRepo repository.ServicePriceRepository, videoProfileRepo repository.VideoProfileRepository) InsertServiceUseCase {
	return &insertServiceUseCase{
		serviceDetailRepo: serviceDetailRepo,
		servicePriceRepo:  servicePriceRepo,
		videoProfileRepo:  videoProfileRepo,
	}
}
