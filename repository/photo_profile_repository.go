package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type PhotoProfileRepository interface {
	Insert(customersPhoto *model.PhotoProfile) error
	FindAllBy(preload string, condition string, searchValue ...interface{}) ([]model.PhotoProfile, error)
	FindById(id int) (model.PhotoProfile, error)
	RetrieveAll(page int, itemPerPage int) ([]model.PhotoProfile, error)
	Update(customersPhoto *model.PhotoProfile, by map[string]interface{}) error
	Delete(id int) error
}

type photoProfileRepository struct {
	db *gorm.DB
}

func (p *photoProfileRepository) Delete(id int) error {
	result := p.db.Delete(&model.PhotoProfile{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (p *photoProfileRepository) Update(customersPhoto *model.PhotoProfile, by map[string]interface{}) error {
	result := p.db.Model(customersPhoto).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (p *photoProfileRepository) RetrieveAll(page int, itemPerPage int) ([]model.PhotoProfile, error) {
	var customersPhotos []model.PhotoProfile
	offset := itemPerPage * (page - 1)
	res := p.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&customersPhotos)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customersPhotos, nil
}

func (p *photoProfileRepository) FindById(id int) (model.PhotoProfile, error) {
	var customersPhoto model.PhotoProfile
	result := p.db.First(&customersPhoto, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customersPhoto, nil
		} else {
			return customersPhoto, err
		}
	}
	return customersPhoto, nil
}

func (p *photoProfileRepository) FindAllBy(preload string, condition string, searchValue ...interface{}) ([]model.PhotoProfile, error) {
	var customersPhotos []model.PhotoProfile
	if preload == "" {
		result := p.db.Where(condition, searchValue...).Find(&customersPhotos)
		if err := result.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			} else {
				return nil, err
			}
		}
	} else {
		result := p.db.Preload(preload).Where(condition, searchValue...).Find(&customersPhotos)
		if err := result.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			} else {
				return nil, err
			}
		}
	}
	return customersPhotos, nil
}

func (p *photoProfileRepository) Insert(customersPhoto *model.PhotoProfile) error {
	result := p.db.Create(customersPhoto)
	return result.Error
}

func NewPhotoProfileRepository(db *gorm.DB) PhotoProfileRepository {
	repo := new(photoProfileRepository)
	repo.db = db
	return repo
}
