package repository

import (
	"errors"
	"surpreedz-backend/model"
	"time"

	"gorm.io/gorm"
)

type SignUpRepository interface {
	SignUpAccount(accountFormInfo *model.AccountFormInfo) error
}

type signUpRepository struct {
	db *gorm.DB
}

func (s *signUpRepository) SignUpAccount(accountFormInfo *model.AccountFormInfo) error {

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	//create account
	newAccount := &model.Account{
		Email:    accountFormInfo.Email,
		Password: accountFormInfo.Password,
		JoinDate: time.Now(),
	}

	if err := tx.Create(newAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	//find id account by username
	var account model.Account
	result := tx.Where("mst_account.email = ?", newAccount.Email).First(&account)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		} else {
			tx.Rollback()
			return err
		}
	}

	//create account_detail
	newAccountDetail := &model.AccountDetail{
		AccountId: account.ID,
		UserName:  accountFormInfo.Name,
		Location:  accountFormInfo.Location,
	}

	if err := tx.Create(newAccountDetail).Error; err != nil {
		tx.Rollback()
		return err
	}

	//find id account_detail by account_id
	var accountDetail model.AccountDetail
	result = tx.Where("mst_account_detail.account_id = ?", newAccountDetail.AccountId).First(&accountDetail)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		} else {
			tx.Rollback()
			return err
		}
	}

	//create photo_profile
	newPhotoProfile := &model.PhotoProfile{
		AccountDetailId: accountDetail.ID,
		PhotoLink:       accountFormInfo.PhotoLink,
		IsDeleted:       accountFormInfo.IsDeleted,
	}

	if err := tx.Create(newPhotoProfile).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func NewSignUpRepository(db *gorm.DB) SignUpRepository {
	repo := new(signUpRepository)
	repo.db = db
	return repo
}
