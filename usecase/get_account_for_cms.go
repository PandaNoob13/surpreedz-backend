package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindAccountCMSUseCase interface {
	FindAccountWithCondition(table string, condition string) ([]model.Account, []string, error)
}

type findAccountCMSUseCase struct {
	cmsAccountRepo repository.CMSAccountRepository
}

func (f *findAccountCMSUseCase) FindAccountWithCondition(table string, condition string) ([]model.Account, []string, error) {
	return f.cmsAccountRepo.GetAllAccount(table, condition)
}

func NewFindAccountCMSUseCase(cmsAccountRepo repository.CMSAccountRepository) FindAccountCMSUseCase {
	return &findAccountCMSUseCase{
		cmsAccountRepo: cmsAccountRepo,
	}
}
