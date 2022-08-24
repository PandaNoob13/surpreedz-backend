package usecase

import (
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
)

type EditAccountUsecase interface {
	EditAccountInfo(accountEditInfo *dto.AccountEditInfo) error
}

type editAccountUsecase struct {
	editAccountRepo repository.EditAccountRepository
}

func (e *editAccountUsecase) EditAccountInfo(accountEditInfo *dto.AccountEditInfo) error {
	return e.editAccountRepo.EditAccount(accountEditInfo)
}

func NewEditAccountUsecase(editAccountRepo repository.EditAccountRepository) EditAccountUsecase {
	return &editAccountUsecase{
		editAccountRepo: editAccountRepo,
	}
}
