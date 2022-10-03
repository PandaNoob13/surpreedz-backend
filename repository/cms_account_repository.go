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

type CMSAccountRepository interface {
	GetAllAccount(table string, condition string) ([]model.Account, []string, error)
}

type cmsAccountRepository struct {
	db  *gorm.DB
	azr *azblob.ServiceClient
}

func (c *cmsAccountRepository) GetAllAccount(table string, condition string) ([]model.Account, []string, error) {
	var customers []model.Account
	result := c.db.Joins("inner join mst_account_detail on mst_account.id = mst_account_detail.account_id")
	result = result.Preload("ServiceDetail.ServicePrices").Preload("ServiceDetail").Preload("AccountDetail").Preload("AccountDetail.PhotoProfiles").Where(table, condition).Find(&customers)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(customers)
			return nil, nil, err
		}
	}


	containerClient, err := c.azr.NewContainerClient("photoprofile")
	if err != nil {
		log.Fatalln("Error getting container client")
	}

	for i := range customers {
		var dataURLArr []string
		blockBlobClient, err := containerClient.NewBlockBlobClient(customers[i].AccountDetail.PhotoProfiles[len(customers[i].AccountDetail.PhotoProfiles)-1].PhotoLink)
		if err != nil {
			fmt.Println(err)
		}
		blobDownloadResponse, err := blockBlobClient.Download(context.TODO(), nil)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("BlobDownloadResponse: ", blobDownloadResponse)
			reader := blobDownloadResponse.Body(nil)
			downloadData, err := io.ReadAll(reader)
			if err != nil {
				fmt.Println(err)
			}
			dataUrl := base64.StdEncoding.EncodeToString(downloadData)
			fmt.Println("Service detail : ", customers[i].ServiceDetail)
			dataURLArr = append(dataURLArr, dataUrl)
			return customers, dataURLArr, nil
		}
	}
	return customers, nil, nil
}

func NewCMSAccountRepository(db *gorm.DB, azr *azblob.ServiceClient) CMSAccountRepository {
	repo := new(cmsAccountRepository)
	repo.db = db
	repo.azr = azr
	return repo
}