package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindBuyerSellerAccCMSUseCase interface {
	FindBuyerAcc() ([]model.Account, []string, error)
	FindSellerAcc() ([]model.Account, []string, error)
	FindBuyerSellerAcc() ([]model.Account, []string, error)
}

type findBuyerSellerAccCMSUseCase struct {
	cmsBuyerSellerAccRepo repository.CMSAccSellerBuyerRepository
}

func (f *findBuyerSellerAccCMSUseCase) FindBuyerAcc() ([]model.Account, []string, error) {
	return f.cmsBuyerSellerAccRepo.GetAllBuyerAcc()
}

func (f *findBuyerSellerAccCMSUseCase) FindSellerAcc() ([]model.Account, []string, error) {
	return f.cmsBuyerSellerAccRepo.GetAllSellerAcc()
}

func (f *findBuyerSellerAccCMSUseCase) FindBuyerSellerAcc() ([]model.Account, []string, error) {
	return f.cmsBuyerSellerAccRepo.GetAllBuyerSellerAcc()
}

func NewFindBuyerSellerAccCMSUseCase(cmsBuyerSellerAccRepo repository.CMSAccSellerBuyerRepository) FindBuyerSellerAccCMSUseCase {
	return &findBuyerSellerAccCMSUseCase{
		cmsBuyerSellerAccRepo: cmsBuyerSellerAccRepo,
	}
}
