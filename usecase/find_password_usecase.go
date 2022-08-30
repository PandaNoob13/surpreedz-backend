package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindPasswordUseCase interface {
	FindPasswordByAccountId(id int) (model.Password, error)
}

type findPasswordUseCase struct {
	passwordRepo repository.PasswordRepository
}

func (f *findPasswordUseCase) FindPasswordByAccountId(id int) (model.Password, error) {
	return f.passwordRepo.FindByAccountId(id)
}

func NewFindPasswordUseCase(passwordRepo repository.PasswordRepository) FindPasswordUseCase {
	return &findPasswordUseCase{
		passwordRepo: passwordRepo,
	}
}
