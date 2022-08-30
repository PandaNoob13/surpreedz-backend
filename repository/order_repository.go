package repository

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *model.Order) error
	FindById(id int) (model.Order, error)
	FindAll(page int, itemPerPage int) ([]model.Order, error)
	FindByBuyerId(id int) (model.Order, error)
	FindAllByBuyerId(buyerId int) ([]dto.AccountCreateDto, error)
	FindAllByServiceDetailId(serviceDetailId int) ([]model.Account, error)
	UpdateByID(order *model.Order, by map[string]interface{}) error
	Delete(order *model.Order) error
}

type orderRepository struct {
	db  *gorm.DB
	azr *azblob.ServiceClient
}

func (o *orderRepository) FindAllByBuyerId(buyerId int) ([]dto.AccountCreateDto, error) {
	var account []model.Account
	result := o.db.Joins("inner join mst_order on mst_account.id = mst_order.buyer_id ")
	result = result.Where("mst_order.buyer_id = ?", buyerId).Distinct("mst_account.id").Preload("AccountDetail").Preload("AccountDetail.PhotoProfiles").Preload("Orders").Preload("Orders.OrderRequest").Preload("Orders.OrderStatus").Preload("ServiceDetail").Preload("ServiceDetail.ServicePrices").Find(&account)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []dto.AccountCreateDto{}, nil
		} else {
			return []dto.AccountCreateDto{}, err
		}
	}
	for accountIndex, accountValue := range account {
		var rightOrders []model.Order
		for _, orderValue := range accountValue.Orders {
			if orderValue.BuyerId == buyerId {
				rightOrders = append(rightOrders, orderValue)
			}
		}
		account[accountIndex].Orders = rightOrders
	}
	containerClient, err := o.azr.NewContainerClient("photoprofile")
	if err != nil {
		log.Fatalln("Error getting container client")
	}
	var tempAccountList []dto.AccountCreateDto
	for index, accountValue := range account {
		var tempAccount dto.AccountCreateDto
		if accountValue.ServiceDetail.SellerId != 0 {
			fmt.Println("Log loop : ", index)
			blockBlobClient, err := containerClient.NewBlockBlobClient(accountValue.AccountDetail.PhotoProfiles[len(accountValue.AccountDetail.PhotoProfiles)-1].PhotoLink)
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
			} else {
				dataUrl := base64.StdEncoding.EncodeToString(downloadData)
				tempAccount.DataUrl = dataUrl
				tempAccount.Account = account[index]
				tempAccountList = append(tempAccountList, tempAccount)
				reader.Close()
			}
		}
		tempAccountList = append(tempAccountList, tempAccount)
	}
	return tempAccountList, nil
}

func (o *orderRepository) FindAllByServiceDetailId(serviceDetailId int) ([]model.Account, error) {
	var account []model.Account
	result := o.db.Joins("inner join mst_order on mst_account.id = mst_order.buyer_id ")
	result = result.Where("mst_order.service_detail_id = ?", serviceDetailId).Distinct("mst_account.id").Preload("AccountDetail").Preload("AccountDetail.PhotoProfiles").Preload("Orders").Preload("Orders.OrderRequest").Preload("Orders.OrderStatus").Preload("ServiceDetail").Preload("ServiceDetail.ServicePrices").Find(&account)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Account{}, nil
		} else {
			return []model.Account{}, err
		}
	}
	for accountIndex, accountValue := range account {
		var rightOrders []model.Order
		for _, orderValue := range accountValue.Orders {
			if orderValue.ServiceDetailId == serviceDetailId {
				rightOrders = append(rightOrders, orderValue)
			}
		}
		account[accountIndex].Orders = rightOrders
	}
	return account, nil
}

// func (o *orderRepository) FindAllByServiceDetailId(serviceDetailId int) ([]model.Order, error) {
// 	var order []model.Order
// 	result := o.db.Where("mst_order.service_detail_id = ?", serviceDetailId).Preload("OrderStatus").Preload("OrderRequest").Preload("ServiceDetail").Preload("ServiceDetail.ServicePrice").Find(&order)
// 	if err := result.Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return order, nil
// 		} else {
// 			return order, err
// 		}
// 	}
// 	return order, nil
// }

func (o *orderRepository) FindByBuyerId(id int) (model.Order, error) {
	var order model.Order
	result := o.db.Where("mst_order.buyer_id = ?", id).Last(&order)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, nil
		} else {
			return order, err
		}
	}
	return order, nil
}

func (o *orderRepository) Create(order *model.Order) error {
	result := o.db.Create(order).Error
	return result
}

func (o *orderRepository) FindById(id int) (model.Order, error) {
	var order model.Order
	result := o.db.Preload("OrderStatus.Refund").Preload("OrderRequest").Preload("Feedback").Preload("VideoResult").Where("mst_order.id = ?", id).First(&order)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, nil
		} else {
			return order, err
		}
	}
	return order, nil
}

func (o *orderRepository) FindAll(page int, itemPerPage int) ([]model.Order, error) {
	var order []model.Order
	offset := itemPerPage * (page - 1)
	result := o.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Preload("OrderStatus.Refund").Preload("OrderRequest").Preload("Feedback").Preload("VideoResult").Find(&order)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, nil
		} else {
			return order, err
		}
	}
	return order, nil
}

func (o *orderRepository) UpdateByID(order *model.Order, by map[string]interface{}) error {
	result := o.db.Model(order).Updates(by).Error
	return result
}

func (o *orderRepository) Delete(order *model.Order) error {
	result := o.db.Delete(order).Error
	return result
}

func NewOrderRepository(db *gorm.DB, azr *azblob.ServiceClient) OrderRepository {
	repo := new(orderRepository)
	repo.db = db
	repo.azr = azr
	return repo
}
