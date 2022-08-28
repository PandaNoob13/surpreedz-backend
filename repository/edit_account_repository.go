package repository

import (
	"errors"
	"fmt"
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/vincent-petithory/dataurl"
	"gorm.io/gorm"
)

type EditAccountRepository interface {
	EditProfile(editProfileDto *dto.EditProfileDto) error
	EditPassword(EditPasswordDto *dto.EditPasswordDto) error
}

type editAccountRepository struct {
	db  *gorm.DB
	azr *azblob.ServiceClient
}

func (e *editAccountRepository) EditProfile(editProfileDto *dto.EditProfileDto) error {
	tx := e.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	//find id account_detail by account_id
	var accountDetail model.AccountDetail
	result := tx.Where("mst_account_detail.account_id = ?", editProfileDto.AccountId).First(&accountDetail)
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
		"name":     editProfileDto.Name,
		"location": editProfileDto.Location,
	}).Error; err != nil {
		fmt.Println(err)
		return err
	}

	// container, err := e.azr.NewContainerClient("photoprofile")
	// if err != nil {
	// 	log.Fatalln("Error getting container client")
	// }
	fmt.Println(editProfileDto.DataUrl)
	dataURL, err := dataurl.DecodeString(editProfileDto.DataUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("content type: %s, data: %s\n", dataURL.MediaType.ContentType(), dataURL.Data)

	//create photo_profile
	newPhotoProfile := &model.PhotoProfile{
		AccountDetailId: accountDetail.ID,
		PhotoLink:       "",
		IsDeleted:       false,
	}

	if err := tx.Create(newPhotoProfile).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Rollback().Error
}

func (e *editAccountRepository) EditPassword(EditPasswordDto *dto.EditPasswordDto) error {
	tx := e.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	//find id password by account_id
	var password model.Password
	result := tx.Where("mst_password.account_id = ?", EditPasswordDto.AccountId).First(&password)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		} else {
			tx.Rollback()
			return err
		}
	}

	//update password
	passwordExist := model.Password{
		ID: password.ID,
	}

	if err := tx.Model(&passwordExist).Updates(map[string]interface{}{
		// "email":    accountEditInfo.Email,
		"password": EditPasswordDto.Password,
	}).Error; err != nil {
		return err
	}
	return tx.Commit().Error
}

func NewEditAccountRepository(db *gorm.DB, azr *azblob.ServiceClient) EditAccountRepository {
	repo := new(editAccountRepository)
	repo.db = db
	repo.azr = azr
	return repo
}
