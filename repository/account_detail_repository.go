package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type AccountDetailRepository interface {
	Insert(customerDetail *model.AccountDetail) error
	FindById(id int) (model.AccountDetail, error)
	RetrieveAll(page int, itemPerPage int) ([]model.AccountDetail, error)
	Update(customerDetail *model.AccountDetail, by map[string]interface{}) error
	Delete(id int) error
}

type accountDetailRepository struct {
	db *gorm.DB
}

func (ad *accountDetailRepository) Delete(id int) error {
	result := ad.db.Delete(&model.AccountDetail{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (ad *accountDetailRepository) Update(customer *model.AccountDetail, by map[string]interface{}) error {
	result := ad.db.Model(customer).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (ad *accountDetailRepository) RetrieveAll(page int, itemPerPage int) ([]model.AccountDetail, error) {
	var customerDetails []model.AccountDetail
	offset := itemPerPage * (page - 1)
	res := ad.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Preload("PhotoProfiles").Find(&customerDetails)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customerDetails, nil
}

func (ad *accountDetailRepository) FindById(id int) (model.AccountDetail, error) {
	var customerDetail model.AccountDetail
	result := ad.db.Preload("PhotoProfiles").Where("mst_account_detail.id = ?", id).First(&customerDetail)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customerDetail, nil
		} else {
			return customerDetail, err
		}
	}
	return customerDetail, nil
}

func (ad *accountDetailRepository) Insert(customerDetail *model.AccountDetail) error {
	result := ad.db.Create(customerDetail)
	return result.Error
}

func NewAccountDetailRepository(db *gorm.DB) AccountDetailRepository {
	repo := new(accountDetailRepository)
	repo.db = db
	return repo
}
