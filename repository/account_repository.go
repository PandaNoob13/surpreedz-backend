package repository

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"surpreedz-backend/model"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Insert(customer *model.Account) error
	FindById(id int) (model.Account, error)
	FindByEmail(email string) (model.Account, string, error)
	RetrieveAll(page int, itemPerPage int) ([]model.Account, error)
	Update(customer *model.Account, by map[string]interface{}) error
	Delete(id int) error
}

type accountRepository struct {
	db  *gorm.DB
	azr *azblob.ServiceClient
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
	res := a.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Preload("Orders.VideoResult").Preload("Orders.Feedback").Preload("Orders.OrderRequest").Preload("Orders.OrderStatus.Refund").Preload("ServiceDetail.VideoProfiles").Preload("ServiceDetail.ServicePrices").Preload("AccountDetail.PhotoProfiles").Find(&customers)
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
	result := a.db.Preload("Orders.VideoResult").Preload("Orders.Feedback").Preload("Orders.OrderRequest").Preload("Orders.OrderStatus.Refund").Preload("ServiceDetail.VideoProfiles").Preload("ServiceDetail.ServicePrices").Preload("AccountDetail.PhotoProfiles").Where("mst_account.id = ?", id).First(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (a *accountRepository) FindByEmail(email string) (model.Account, string, error) {
	var customer model.Account
	result := a.db.Preload("ServiceDetail.ServicePrices").Preload("ServiceDetail").Preload("AccountDetail.PhotoProfiles").Where("mst_account.email = ?", email).First(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(customer)
			return customer, "", err
		}
	}
	containerClient, err := a.azr.NewContainerClient("photoprofile")
	if err != nil {
		log.Fatalln("Error getting container client")
	}
	blockBlobClient, err := containerClient.NewBlockBlobClient(customer.AccountDetail.PhotoProfiles[len(customer.AccountDetail.PhotoProfiles)-1].PhotoLink)
	if err != nil {
		fmt.Println(err)
	}
	blobDownloadResponse, err := blockBlobClient.Download(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
	}
	reader := blobDownloadResponse.Body(nil)
	downloadData, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	dataUrl := base64.StdEncoding.EncodeToString(downloadData)
	fmt.Println("Service detail : ", customer.ServiceDetail)
	return customer, dataUrl, nil
}

func (a *accountRepository) Insert(customer *model.Account) error {
	result := a.db.Create(customer)
	return result.Error
}

func NewAccountRepository(db *gorm.DB, azr *azblob.ServiceClient) AccountRepository {
	repo := new(accountRepository)
	repo.db = db
	repo.azr = azr
	return repo
}
