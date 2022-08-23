package manager

import "surpreedz-backend/usecase"

type UseCaseManager interface {
	AddService() usecase.InsertServiceUseCase
	AddOrder() usecase.InsertOrderUseCase
	AddOrderStatus() usecase.InsertOrderStatusUseCase
	UpdateService() usecase.UpdateServiceUseCase
	FindService() usecase.FindServiceUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) AddOrderStatus() usecase.InsertOrderStatusUseCase {
	return usecase.NewInsertOrderStatusUseCase(u.repoManager.OrderStatusRepo(), u.repoManager.RefundRepository())
}

func (u *useCaseManager) AddOrder() usecase.InsertOrderUseCase {
	return usecase.NewInsertOrderUseCase(u.repoManager.OrderRepo(), u.repoManager.OrderRequestRepo())
}

func (u *useCaseManager) AddService() usecase.InsertServiceUseCase {
	return usecase.NewInsertServiceUseCase(u.repoManager.AccountDetailRepo(), u.repoManager.ServiceDetailRepo(), u.repoManager.ServicePriceRepo(), u.repoManager.VideoProfileRepo())
}

func (u *useCaseManager) UpdateService() usecase.UpdateServiceUseCase {
	return usecase.NewUpdateServiceUseCase(u.repoManager.ServiceDetailRepo())
}

func (u *useCaseManager) FindService() usecase.FindServiceUseCase {
	return usecase.NewFindServiceUseCase(u.repoManager.ServiceDetailRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
