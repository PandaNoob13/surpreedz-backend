package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindPasswordUseCase interface {
	FindPasswordById(id int) (model.Password, error)
}

type findPasswordUseCase struct {
	passwordRepo repository.PasswordRepository
}

func (f *findPasswordUseCase) FindPasswordById(id int) (model.Password, error) {
	return f.passwordRepo.FindById(id)
}

func NewFindPasswordUseCase(passwordRepo repository.PasswordRepository) FindPasswordUseCase {
	return &findPasswordUseCase{
		passwordRepo: passwordRepo,
	}
}
