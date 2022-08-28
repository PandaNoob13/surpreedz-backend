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

type ServiceDetailRepository interface {
	Insert(customersService *model.ServiceDetail) error
	FindById(id int) (model.ServiceDetail, error)
	FindBySellerId(id int) (model.ServiceDetail, error)
	RetrieveAll(page int, itemPerPage int) ([]model.ServiceDetail, error)
	HomePageRetrieveAll(page int, itemPerPage int) ([]model.Account, error)
	Update(customersService *model.ServiceDetail, by map[string]interface{}) error
	Delete(id int) error
}

type serviceDetailRepository struct {
	db  *gorm.DB
	azr *azblob.ServiceClient
}

func (s *serviceDetailRepository) Delete(id int) error {
	result := s.db.Delete(&model.ServiceDetail{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (s *serviceDetailRepository) Update(customersService *model.ServiceDetail, by map[string]interface{}) error {
	result := s.db.Model(customersService).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (s *serviceDetailRepository) RetrieveAll(page int, itemPerPage int) ([]model.ServiceDetail, error) {
	var customersServices []model.ServiceDetail
	offset := itemPerPage * (page - 1)
	res := s.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Preload("Orders.VideoResult").Preload("Orders.Feedback").Preload("Orders.OrderRequest").Preload("Orders.OrderStatus.Refund").Preload("ServicePrices").Preload("VideoProfiles").Find(&customersServices)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customersServices, nil
}

func (s *serviceDetailRepository) HomePageRetrieveAll(page int, itemPerPage int) ([]model.Account, error) {
	var homepageServices []model.Account
	offset := itemPerPage * (page - 1)
	res := s.db.Order("created_at").Limit(itemPerPage).Offset(offset).Preload("AccountDetail").Preload("AccountDetail.PhotoProfiles").Preload("ServiceDetail").Preload("ServiceDetail.VideoProfiles").Preload("ServiceDetail.ServicePrices").Find(&homepageServices)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	containerClient, err := s.azr.NewContainerClient("photoprofile")
	if err != nil {
		log.Fatalln("Error getting container client")
	}
	var homePageRetrieval []model.Account
	for index, hp := range homepageServices {
		if hp.ServiceDetail.SellerId != 0 {
			fmt.Println("Log loop : ", index)
			blockBlobClient, err := containerClient.NewBlockBlobClient(hp.AccountDetail.PhotoProfiles[len(hp.AccountDetail.PhotoProfiles)-1].PhotoLink)
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
			homepageServices[index].DataUrl = dataUrl
			homepageServices[index].StringJoinDate = homepageServices[index].JoinDate.Format("2006-January-02")
			homePageRetrieval = append(homePageRetrieval, homepageServices[index])
			reader.Close()
		}
	}
	return homePageRetrieval, nil
}

func (s *serviceDetailRepository) FindById(id int) (model.ServiceDetail, error) {
	var customersService model.ServiceDetail
	result := s.db.Preload("Orders.VideoResult").Preload("Orders.Feedback").Preload("Orders.OrderRequest").Preload("Orders.OrderStatus.Refund").Preload("ServicePrices").Preload("VideoProfiles").Where("mst_service_detail.id = ?", id).First(&customersService)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customersService, nil
		} else {
			return customersService, err
		}
	}
	return customersService, nil
}

func (s *serviceDetailRepository) FindBySellerId(id int) (model.ServiceDetail, error) {
	var customersService model.ServiceDetail
	result := s.db.Where("mst_service_detail.seller_id = ?", id).Last(&customersService)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customersService, nil
		} else {
			return customersService, err
		}
	}
	return customersService, nil
}

func (s *serviceDetailRepository) Insert(customersService *model.ServiceDetail) error {
	result := s.db.Create(customersService)
	return result.Error
}

func NewServiceDetailRepository(db *gorm.DB, azr *azblob.ServiceClient) ServiceDetailRepository {
	repo := new(serviceDetailRepository)
	repo.db = db
	repo.azr = azr
	return repo
}
