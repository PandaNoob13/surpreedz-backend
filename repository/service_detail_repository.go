package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type ServiceDetailRepository interface {
	Insert(customersService *model.ServiceDetail) error
	FindAllBy(preload string, condition string, serachValue ...interface{}) ([]model.ServiceDetail, error)
	FindById(id int) (model.ServiceDetail, error)
	RetrieveAll(page int, itemPerPage int) ([]model.ServiceDetail, error)
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
	res := s.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&customersServices)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customersServices, nil
}

func (s *serviceDetailRepository) FindById(id int) (model.ServiceDetail, error) {
	var customersService model.ServiceDetail
	result := s.db.First(&customersService, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customersService, nil
		} else {
			return customersService, err
		}
	}
	return customersService, nil
}

func (s *serviceDetailRepository) FindAllBy(preload string, condition string, searchValue ...interface{}) ([]model.ServiceDetail, error) {
	var customersServices []model.ServiceDetail
	if preload == "" {
		result := s.db.Where(condition, searchValue...).Find(&customersServices)
		if err := result.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			} else {
				return nil, err
			}
		}
	} else {
		result := s.db.Preload(preload).Where(condition, searchValue...).Find(&customersServices)
		if err := result.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			} else {
				return nil, err
			}
		}
	}
	return customersServices, nil
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
