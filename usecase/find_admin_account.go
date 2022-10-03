package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindAdminAccUseCase interface {
	FindAdminAccByUsername(uname string) (model.Admin, error)
}

type findAdminAccUseCase struct {
	adminAccRepo repository.AdminAccRepository
}

func (f *findAdminAccUseCase) FindAdminAccByUsername(uname string) (model.Admin, error) {
	return f.adminAccRepo.FindByUsername(uname)
}

func NewFindAdminAccUseCase(adminAccRepo repository.AdminAccRepository) FindAdminAccUseCase {
	return &findAdminAccUseCase{
		adminAccRepo: adminAccRepo,
	}
}
