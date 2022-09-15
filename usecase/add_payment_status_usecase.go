package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type AddPaymentStatusUseCase interface {
	AddPaymentStatus(paymentStatus *model.PaymentStatus) error
}

type addPaymentStatusUseCase struct {
	paymentStatusRepo repository.PaymentStatusRepository
}

func (a *addPaymentStatusUseCase) AddPaymentStatus(paymentStatus *model.PaymentStatus) error {
	return a.paymentStatusRepo.Insert(paymentStatus)
}

func NewAddPaymentStatusUseCase(paymentStatusRepo repository.PaymentStatusRepository) AddPaymentStatusUseCase {
	return &addPaymentStatusUseCase{
		paymentStatusRepo: paymentStatusRepo,
	}
}
