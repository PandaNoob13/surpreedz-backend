package manager

import "surpreedz-backend/usecase"

type UseCaseManager interface {
	AddOrder() usecase.InsertOrderUseCase
	AddOrderStatus() usecase.InsertOrderStatusUseCase

	AddService() usecase.InsertServiceUseCase
	UpdateService() usecase.UpdateServiceUseCase
	FindService() usecase.FindServiceUseCase
	RetrieveServiceHomePage() usecase.ShowServicesHomePageUseCase

	AddVideoResult() usecase.AddVideoResultUseCase
	RetrieveAllVideoResult() usecase.RetrieveAllVideoResultUseCase
	FindVideoResultById() usecase.FindVideoResultByOrderIdUseCase

	RetrieveAllOrder() usecase.RetrieveAllOrderUseCase
	FindOrderById() usecase.FindOrderByIdUseCase

	AddFeedback() usecase.AddFeedbackUseCase
	FindFeedbackById() usecase.FindFeedbackByIdUseCase
	RetrieveAllFeedback() usecase.RetrieveAllFeedbackUseCase

	FindAccountUseCase() usecase.FindAccountUseCase
	SignUpAccountUseCase() usecase.SignUpUsecase
	EditAccountInfoUsecase() usecase.EditAccountUsecase

	FindPasswordByAccId() usecase.FindPasswordUseCase

	AddPaymentStatus() usecase.AddPaymentStatusUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) AddPaymentStatus() usecase.AddPaymentStatusUseCase {
	return usecase.NewAddPaymentStatusUseCase(u.repoManager.PaymentStatusRepo())
}

func (u *useCaseManager) AddOrderStatus() usecase.InsertOrderStatusUseCase {
	return usecase.NewInsertOrderStatusUseCase(u.repoManager.AddOrderStatusRepo())
}

func (u *useCaseManager) AddOrder() usecase.InsertOrderUseCase {
	return usecase.NewInsertOrderUseCase(u.repoManager.AddOrderRepo())
}

func (u *useCaseManager) AddService() usecase.InsertServiceUseCase {
	return usecase.NewInsertServiceUseCase(u.repoManager.AddServiceRepo())
}

func (u *useCaseManager) UpdateService() usecase.UpdateServiceUseCase {
	return usecase.NewUpdateServiceUseCase(u.repoManager.EditServiceRepo())
}

func (u *useCaseManager) FindService() usecase.FindServiceUseCase {
	return usecase.NewFindServiceUseCase(u.repoManager.ServiceDetailRepo())
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

func (u *useCaseManager) RetrieveServiceHomePage() usecase.ShowServicesHomePageUseCase {
	return usecase.NewShowServiceHomePageUseCase(u.repoManager.ServiceDetailRepo())
}

func (u *useCaseManager) AddVideoResult() usecase.AddVideoResultUseCase {
	return usecase.NewAddVideoResultUseCase(u.repoManager.VideoResultRepo())
}

func (u *useCaseManager) RetrieveAllVideoResult() usecase.RetrieveAllVideoResultUseCase {
	return usecase.NewRetrieveAllVideoResult(u.repoManager.VideoResultRepo())
}

func (u *useCaseManager) FindVideoResultById() usecase.FindVideoResultByOrderIdUseCase {
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

func (u *useCaseManager) FindPasswordByAccId() usecase.FindPasswordUseCase {
	return usecase.NewFindPasswordUseCase(u.repoManager.PasswordRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
