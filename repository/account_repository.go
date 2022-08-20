package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Insert(customer *model.Account) error
	FindAllBy(preload string, condition string, searchValue ...interface{}) ([]model.Account, error)
	FindById(id int) (model.Account, error)
	RetrieveAll(page int, itemPerPage int) ([]model.Account, error)
	Update(customer *model.Account, by map[string]interface{}) error
	Delete(id int) error
}

type accountRepository struct {
	db *gorm.DB
}

func (a *accountRepository) Delete(id int) error {
	result := a.db.Delete(&model.Account{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (a *accountRepository) Update(customer *model.Account, by map[string]interface{}) error {
	result := a.db.Model(customer).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (a *accountRepository) RetrieveAll(page int, itemPerPage int) ([]model.Account, error) {
	var customers []model.Account
	offset := itemPerPage * (page - 1)
	res := a.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&customers)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customers, nil
}

func (a *accountRepository) FindById(id int) (model.Account, error) {
	var customer model.Account
	result := a.db.First(&customer, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (a *accountRepository) FindAllBy(preload string, condition string, searchValue ...interface{}) ([]model.Account, error) {
	var customers []model.Account
	if preload == "" {
		result := a.db.Where(condition, searchValue...).Find(&customers)
		if err := result.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			} else {
				return nil, err
			}
		}
	} else {
		result := a.db.Preload(preload).Where(condition, searchValue...).Find(&customers)
		if err := result.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			} else {
				return nil, err
			}
		}
	}
	return customers, nil
}

func (a *accountRepository) Insert(customer *model.Account) error {
	result := a.db.Create(customer)
	return result.Error
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	repo := new(accountRepository)
	repo.db = db
	return repo
}
