package manager

import "surpreedz-backend/usecase"

type UseCaseManager interface {
	FindAccountUseCase() usecase.FindAccountUseCase
	SignUpAccountUseCase() usecase.SignUpUsecase
	EditAccountInfoUsecase() usecase.EditAccountUsecase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) FindAccountUseCase() usecase.FindAccountUseCase {
	return usecase.NewFindAccountUseCase(u.repoManager.AccountRepo())
}

func (u *useCaseManager) SignUpAccountUseCase() usecase.SignUpUsecase {
	return usecase.NewSignUpUsecase(u.repoManager.SignUpAccountRepo())
}

func (u *useCaseManager) EditAccountInfoUsecase() usecase.EditAccountUsecase {
	return usecase.NewEditAccountUsecase(u.repoManager.EditAccountRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
