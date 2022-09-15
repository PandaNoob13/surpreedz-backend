package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type PaymentStatusRepository interface {
	Insert(paymentStatus *model.PaymentStatus) error
	FindById(id int) (model.PaymentStatus, error)
	FindAll(page int, itemPerPage int) ([]model.PaymentStatus, error)
	Update(paymentStatus *model.PaymentStatus, by []map[string]interface{}) error
	Delete(id int) error
}

type paymentStatusRepository struct {
	db *gorm.DB
}

func (p *paymentStatusRepository) Delete(id int) error {
	result := p.db.Delete(&model.PaymentStatus{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (p *paymentStatusRepository) Update(paymentStatus *model.PaymentStatus, by []map[string]interface{}) error {
	result := p.db.Model(paymentStatus).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (p *paymentStatusRepository) FindAll(page int, itemPerPage int) ([]model.PaymentStatus, error) {
	var paymentStatuses []model.PaymentStatus
	offset := itemPerPage * (page - 1)
	res := p.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&paymentStatuses)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return paymentStatuses, nil
}

func (p *paymentStatusRepository) FindById(id int) (model.PaymentStatus, error) {
	var paymentStatus model.PaymentStatus
	result := p.db.Where("mst_payment_status.id = ?", id).First(&paymentStatus)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return paymentStatus, nil
		} else {
			return paymentStatus, err
		}
	}
	return paymentStatus, nil
}

func (p *paymentStatusRepository) Insert(paymentStatus *model.PaymentStatus) error {
	result := p.db.Create(paymentStatus)
	return result.Error
}

func NewPaymentRepository(db *gorm.DB) PaymentStatusRepository {
	repo := new(paymentStatusRepository)
	repo.db = db
	return repo
}
