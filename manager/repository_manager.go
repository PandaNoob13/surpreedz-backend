package manager

import "surpreedz-backend/repository"

type RepositoryManager interface {
	AccountRepo() repository.AccountRepository
	AccountDetailRepo() repository.AccountDetailRepository
	PhotoProfileRepo() repository.PhotoProfileRepository
	ServiceDetailRepo() repository.ServiceDetailRepository
	ServicePriceRepo() repository.ServicePriceRepository
	VideoProfileRepo() repository.VideoProfileRepository
	FeedbackRepo() repository.FeedbackRepository
	OrderRepo() repository.OrderRepository
	OrderRequestRepo() repository.OrderRequestRepository
	OrderStatusRepo() repository.OrderStatusRepository
	RefundRepository() repository.RefundRepository
	VideoResultRepo() repository.VideoResultRepository
	SignUpAccountRepo() repository.SignUpRepository
	EditAccountRepo() repository.EditAccountRepository
	EditServiceRepo() repository.EditServiceRepository
	AddOrderRepo() repository.AddOrderRepository
	AddOrderStatusRepo() repository.AddOrderStatusRepository
	AddServiceRepo() repository.AddServiceRepository
	PasswordRepo() repository.PasswordRepository
	PaymentStatusRepo() repository.PaymentStatusRepository
	AccountCMSRepo() repository.CMSAccountRepository
	AccountBuyerSellerCMSRepo() repository.CMSAccSellerBuyerRepository
	OrderCMSRepo() repository.CMSOrderRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) PaymentStatusRepo() repository.PaymentStatusRepository {
	return repository.NewPaymentRepository(r.infra.SqlDb())
}

func (r *repositoryManager) AddServiceRepo() repository.AddServiceRepository {
	return repository.NewAddServiceRepository(r.infra.SqlDb())
}

func (r *repositoryManager) AddOrderStatusRepo() repository.AddOrderStatusRepository {
	return repository.NewAddOrderStatusRepository(r.infra.SqlDb())
}

func (r *repositoryManager) AddOrderRepo() repository.AddOrderRepository {
	return repository.NewAddOrderRepository(r.infra.SqlDb())
}

func (r *repositoryManager) AccountRepo() repository.AccountRepository {
	return repository.NewAccountRepository(r.infra.SqlDb(), r.infra.AzrClient())
}

func (r *repositoryManager) AccountDetailRepo() repository.AccountDetailRepository {
	return repository.NewAccountDetailRepository(r.infra.SqlDb())
}

func (r *repositoryManager) PhotoProfileRepo() repository.PhotoProfileRepository {
	return repository.NewPhotoProfileRepository(r.infra.SqlDb())
}

func (r *repositoryManager) ServiceDetailRepo() repository.ServiceDetailRepository {
	return repository.NewServiceDetailRepository(r.infra.SqlDb(), r.infra.AzrClient())
}

func (r *repositoryManager) ServicePriceRepo() repository.ServicePriceRepository {
	return repository.NewServicePriceRepository(r.infra.SqlDb())
}

func (r *repositoryManager) VideoProfileRepo() repository.VideoProfileRepository {
	return repository.NewVideoProfileRepository(r.infra.SqlDb())
}

func (r *repositoryManager) FeedbackRepo() repository.FeedbackRepository {
	return repository.NewFeedbackRepository(r.infra.SqlDb())
}

func (r *repositoryManager) OrderRepo() repository.OrderRepository {
	return repository.NewOrderRepository(r.infra.SqlDb(), r.infra.AzrClient())
}

func (r *repositoryManager) OrderRequestRepo() repository.OrderRequestRepository {
	return repository.NewOrderRequestRepository(r.infra.SqlDb())
}

func (r *repositoryManager) OrderStatusRepo() repository.OrderStatusRepository {
	return repository.NewOrderStatusRepository(r.infra.SqlDb())
}

func (r *repositoryManager) RefundRepository() repository.RefundRepository {
	return repository.NewRefundRepository(r.infra.SqlDb())
}

func (r *repositoryManager) VideoResultRepo() repository.VideoResultRepository {
	return repository.NewVideoResultRepository(r.infra.SqlDb(), r.infra.AzrClient())
}

func (r *repositoryManager) SignUpAccountRepo() repository.SignUpRepository {
	return repository.NewSignUpRepository(r.infra.SqlDb())
}

func (r *repositoryManager) EditAccountRepo() repository.EditAccountRepository {
	return repository.NewEditAccountRepository(r.infra.SqlDb(), r.infra.AzrClient())
}

func (r *repositoryManager) EditServiceRepo() repository.EditServiceRepository {
	return repository.NewEditServiceRepository(r.infra.SqlDb())
}

func (r *repositoryManager) PasswordRepo() repository.PasswordRepository {
	return repository.NewPasswordRepository(r.infra.SqlDb())
}

func (r *repositoryManager) AccountCMSRepo() repository.CMSAccountRepository {
	return repository.NewCMSAccountRepository(r.infra.SqlDb(), r.infra.AzrClient())
}

func (r *repositoryManager) AccountBuyerSellerCMSRepo() repository.CMSAccSellerBuyerRepository {
	return repository.NewCMSAccSellerBuyerRepository(r.infra.SqlDb(), r.infra.AzrClient())
}

func (r *repositoryManager) OrderCMSRepo() repository.CMSOrderRepository {
	return repository.NewCMSOrderRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
