package manager

import "surpreedz-backend/usecase"

type UseCaseManager interface {
	FindAccountUseCase() usecase.FindAccountUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) FindAccountUseCase() usecase.FindAccountUseCase {
	return usecase.NewFindAccountUseCase(u.repoManager.AccountRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
