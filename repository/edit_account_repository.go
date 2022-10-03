package repository

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/utils"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type EditAccountRepository interface {
	EditProfile(editProfileDto *dto.EditProfileDto) error
	EditPassword(EditPasswordDto *dto.EditPasswordDto) error
	EditVerifiedStatus(VerifyFromCMSDto *dto.VerifyFromCMS) error
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

	if editProfileDto.DataUrl != "" {
		containerClient, err := e.azr.NewContainerClient("photoprofile")
		if err != nil {
			log.Fatalln("Error getting container client")
		}

		uid, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
		}
		splittedUrl := strings.Split(editProfileDto.DataUrl, ",")
		//contentType := splittedUrl[0]
		dataUrl := splittedUrl[1]
		image, err := base64.StdEncoding.DecodeString(dataUrl)
		if err != nil {
			fmt.Println(err)
		}
		blockBlobClient, err := containerClient.NewBlockBlobClient(time.Now().Format("20060102") + uid.String() + ".jpg")
		if err != nil {
			fmt.Println(err)
		}

		blobUploadResponse, err := blockBlobClient.UploadBuffer(context.TODO(), image, azblob.UploadOption{})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Upload Response : ", blobUploadResponse)

		//create photo_profile
		newPhotoProfile := &model.PhotoProfile{
			AccountDetailId: accountDetail.ID,
			PhotoLink:       time.Now().Format("20060102") + uid.String() + ".jpg",
			IsDeleted:       false,
		}

		if err := tx.Create(newPhotoProfile).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
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
	passHash, _ := utils.HashPassword(EditPasswordDto.Password)

	passwordExist := model.Password{
		ID: password.ID,
	}

	if err := tx.Model(&passwordExist).Updates(map[string]interface{}{
		// "email":    accountEditInfo.Email,
		"password": passHash,
	}).Error; err != nil {
		return err
	}
	return tx.Commit().Error
}

func (e *editAccountRepository) EditVerifiedStatus(VerifyFromCMSDto *dto.VerifyFromCMS) error {
	var accountDetail model.AccountDetail
	result := e.db.Where("mst_account_detail.account_id = ?", VerifyFromCMSDto.AccountId).First(&accountDetail)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		} else {
			return err
		}
	}

	accountDetailExist := model.AccountDetail{
		ID: accountDetail.ID,
	}

	if err := e.db.Model(&accountDetailExist).Updates(map[string]interface{}{
		"verified_status":  VerifyFromCMSDto.VerifiedStatus,
		"verified_request": VerifyFromCMSDto.VerifiedRequest,
	}).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func NewEditAccountRepository(db *gorm.DB, azr *azblob.ServiceClient) EditAccountRepository {
	repo := new(editAccountRepository)
	repo.db = db
	repo.azr = azr
	return repo
}
