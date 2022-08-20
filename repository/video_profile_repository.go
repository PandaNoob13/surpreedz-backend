package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type VideoProfileRepository interface {
	Insert(serviceVideo *model.VideoProfile) error
	FindById(id int) (model.VideoProfile, error)
	RetrieveAll(page int, itemPerPage int) ([]model.VideoProfile, error)
	Update(serviceVideo *model.VideoProfile, by map[string]interface{}) error
	Delete(id int) error
}

type videoProfileRepository struct {
	db *gorm.DB
}

func (v *videoProfileRepository) Delete(id int) error {
	result := v.db.Delete(&model.VideoProfile{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (v *videoProfileRepository) Update(serviceVideo *model.VideoProfile, by map[string]interface{}) error {
	result := v.db.Model(serviceVideo).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (v *videoProfileRepository) RetrieveAll(page int, itemPerPage int) ([]model.VideoProfile, error) {
	var serviceVideos []model.VideoProfile
	offset := itemPerPage * (page - 1)
	res := v.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&serviceVideos)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return serviceVideos, nil
}

func (v *videoProfileRepository) FindById(id int) (model.VideoProfile, error) {
	var serviceVideo model.VideoProfile
	result := v.db.First(&serviceVideo, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serviceVideo, nil
		} else {
			return serviceVideo, err
		}
	}
	return serviceVideo, nil
}

func (v *videoProfileRepository) Insert(serviceVideo *model.VideoProfile) error {
	result := v.db.Create(serviceVideo)
	return result.Error
}

func NewVideoProfileRepository(db *gorm.DB) VideoProfileRepository {
	repo := new(videoProfileRepository)
	repo.db = db
	return repo
}
