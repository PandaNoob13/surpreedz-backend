package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type RefundRepository interface {
	Create(refund *model.Refund) error
	FindById(id int) (model.Refund, error)
	FindAll() ([]model.Refund, error)
	UpdateByID(refund *model.Refund, by map[string]interface{}) error
	Delete(refund *model.Refund) error
}

type refundRepository struct {
	db *gorm.DB
}

func (r *refundRepository) Create(refund *model.Refund) error {
	result := r.db.Create(refund).Error
	return result
}

func (r *refundRepository) FindById(id int) (model.Refund, error) {
	var refund model.Refund
	result := r.db.First(&refund, "order_status_id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return refund, nil
		} else {
			return refund, err
		}
	}
	return refund, nil
}

func (r *refundRepository) FindAll() ([]model.Refund, error) {
	var refunds []model.Refund
	result := r.db.Find(&refunds)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return refunds, nil
		} else {
			return refunds, err
		}
	}
	return refunds, nil
}

func (r *refundRepository) UpdateByID(refund *model.Refund, by map[string]interface{}) error {
	result := r.db.Model(refund).Updates(by).Error
	return result
}

func (r *refundRepository) Delete(refund *model.Refund) error {
	result := r.db.Delete(refund).Error
	return result
}

func NewRefundRepository(db *gorm.DB) RefundRepository {
	repo := new(refundRepository)
	repo.db = db
	return repo
}
