package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type EditAccountRepository interface {
	EditAccount(AccountEditInfo *model.AccountEditInfo) error
}

type editAccountRepository struct {
	db *gorm.DB
}

func (e *editAccountRepository) EditAccount(accountEditInfo *model.AccountEditInfo) error {

	tx := e.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	//update account
	accountExist := model.Account{
		ID: accountEditInfo.ID,
	}

	if err := tx.Model(&accountExist).Updates(map[string]interface{}{
		"email":    accountEditInfo.Email,
		"password": accountEditInfo.Password,
	}).Error; err != nil {
		return err
	}

	//find id account_detail by account_id
	var accountDetail model.AccountDetail
	result := tx.Where("mst_account_detail.account_id = ?", accountEditInfo.ID).First(&accountDetail)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		} else {
			tx.Rollback()
			return err
		}
	}

	//update account detail
	accountDetailExist := model.AccountDetail{
		ID: accountDetail.ID,
	}

	if err := tx.Model(&accountDetailExist).Updates(map[string]interface{}{
		"user_name":    accountEditInfo.Name,
		"location": accountEditInfo.Location,
	}).Error; err != nil {
		return err
	}

	//create photo_profile
	newPhotoProfile := &model.PhotoProfile{
		AccountDetailId: accountDetail.ID,
		PhotoLink:       accountEditInfo.PhotoLink,
		IsDeleted:       accountEditInfo.IsDeleted,
	}

	if err := tx.Create(newPhotoProfile).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func NewEditAccountRepository(db *gorm.DB) EditAccountRepository {
	repo := new(editAccountRepository)
	repo.db = db
	return repo
}
