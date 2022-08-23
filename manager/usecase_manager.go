package manager

import "surpreedz-backend/usecase"

type UseCaseManager interface {
	AddService() usecase.InsertServiceUseCase
	AddOrder() usecase.InsertOrderUseCase
	AddOrderStatus() usecase.InsertOrderStatusUseCase
	UpdateService() usecase.UpdateServiceUseCase
	FindService() usecase.FindServiceUseCase
	RetrieveServiceHomePage() usecase.ShowServicesHomePageUseCase
	AddVideoResult() usecase.AddVideoResultUseCase
	RetrieveAllVideoResult() usecase.RetrieveAllVideoResultUseCase
	FindVideoResultById() usecase.FindVideoResultByIdUseCase
	RetrieveAllOrder() usecase.RetrieveAllOrderUseCase
	FindOrderById() usecase.FindOrderByIdUseCase
	AddFeedback() usecase.AddFeedbackUseCase
	FindFeedbackById() usecase.FindFeedbackByIdUseCase
	RetrieveAllFeedback() usecase.RetrieveAllFeedbackUseCase
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

func (u *useCaseManager) RetrieveServiceHomePage() usecase.ShowServicesHomePageUseCase {
	return usecase.NewShowServiceHomePageUseCase(u.repoManager.ServiceDetailRepo())
}

func (u *useCaseManager) AddVideoResult() usecase.AddVideoResultUseCase {
	return usecase.NewAddVideoResultUseCase(u.repoManager.VideoResultRepo())
}

func (u *useCaseManager) RetrieveAllVideoResult() usecase.RetrieveAllVideoResultUseCase {
	return usecase.NewRetrieveAllVideoResult(u.repoManager.VideoResultRepo())
}

func (u *useCaseManager) FindVideoResultById() usecase.FindVideoResultByIdUseCase {
	return usecase.NewFindVideoResultByIdUseCase(u.repoManager.VideoResultRepo())
}

func (u *useCaseManager) RetrieveAllOrder() usecase.RetrieveAllOrderUseCase {
	return usecase.NewRetrieveAllOrderUseCas(u.repoManager.OrderRepo())
}

func (u *useCaseManager) FindOrderById() usecase.FindOrderByIdUseCase {
	return usecase.NewFindOrderByIdUseCase(u.repoManager.OrderRepo())
}

func (u *useCaseManager) AddFeedback() usecase.AddFeedbackUseCase {
	return usecase.NewAddFeedbackUseCase(u.repoManager.FeedbackRepo())
}

func (u *useCaseManager) FindFeedbackById() usecase.FindFeedbackByIdUseCase {
	return usecase.NewFindFeedbackByIdUseCase(u.repoManager.FeedbackRepo())
}

func (u *useCaseManager) RetrieveAllFeedback() usecase.RetrieveAllFeedbackUseCase {
	return usecase.NewRetrieveAllFeedbackUseCase(u.repoManager.FeedbackRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
