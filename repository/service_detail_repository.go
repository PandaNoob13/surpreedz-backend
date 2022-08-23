package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type ServiceDetailRepository interface {
	Insert(customersService *model.ServiceDetail) error
	FindById(id int) (model.ServiceDetail, error)
	RetrieveAll(page int, itemPerPage int) ([]model.ServiceDetail, error)
	HomePageRetrieveAll(page int, itemPerPage int) ([]model.ServiceDetail, error)
	Update(customersService *model.ServiceDetail, by map[string]interface{}) error
	Delete(id int) error
}

type serviceDetailRepository struct {
	db *gorm.DB
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

func (s *serviceDetailRepository) HomePageRetrieveAll(page int, itemPerPage int) ([]model.ServiceDetail, error) {
	var homepageServices []model.ServiceDetail
	offset := itemPerPage * (page - 1)
	res := s.db.Order("created_at").Limit(itemPerPage).Offset(offset).Preload("ServicePrices").Preload("VideoProfiles").Find(&homepageServices)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return homepageServices, nil
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

func (s *serviceDetailRepository) Insert(customersService *model.ServiceDetail) error {
	result := s.db.Create(customersService)
	return result.Error
}

func NewServiceDetailRepository(db *gorm.DB) ServiceDetailRepository {
	repo := new(serviceDetailRepository)
	repo.db = db
	return repo
}
