package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
	"surpreedz-backend/utils"
)

type InsertServiceUseCase interface {
	AddService(accountId int, role string, description string, price int, videoLink string) error
}

type insertServiceUseCase struct {
	serviceDetailRepo repository.ServiceDetailRepository
	servicePriceRepo  repository.ServicePriceRepository
	videoProfileRepo  repository.VideoProfileRepository
}

func (s *insertServiceUseCase) AddService(accountId int, role string, description string, price int, videoLink string) error {
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

	serviceDetail, err := s.serviceDetailRepo.FindBySellerId(accountId)
	utils.IsError(err)

	toServicePrice := model.ServicePrice{
		ServiceDetailId: serviceDetail.ID,
		Price:           insertService.Price,
	}
	err2 := s.servicePriceRepo.Insert(&toServicePrice)
	utils.IsError(err2)

	toVideoProfile := model.VideoProfile{
		ServiceDetailId:  serviceDetail.ID,
		VideoProfileLink: "",
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
