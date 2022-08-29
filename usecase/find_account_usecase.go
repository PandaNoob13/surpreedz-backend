package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindAccountUseCase interface {
	FindAccountByEmail(email string) (model.Account, string, error)
}

type findAccountUseCase struct {
	accountRepo repository.AccountRepository
}

func (a *findAccountUseCase) FindAccountByEmail(email string) (model.Account, string, error) {
	return a.accountRepo.FindByEmail(email)
}

func NewFindAccountUseCase(accountRepo repository.AccountRepository) FindAccountUseCase {
	return &findAccountUseCase{
		accountRepo: accountRepo,
	}
}
