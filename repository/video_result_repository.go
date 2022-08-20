package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type VideoResultRepository interface {
	Create(videoResult *model.VideoResult) error
	FindById(id int) (model.VideoResult, error)
	FindAll() ([]model.VideoResult, error)
	UpdateByID(videoResult *model.VideoResult, by map[string]interface{}) error
	Delete(videoResult *model.VideoResult) error
}

type videoResultRepository struct {
	db *gorm.DB
}

func (v *videoResultRepository) Create(videoResult *model.VideoResult) error {
	result := v.db.Create(videoResult).Error
	return result
}

func (v *videoResultRepository) FindById(id int) (model.VideoResult, error) {
	var videoResult model.VideoResult
	result := v.db.First(&videoResult, "order_id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return videoResult, nil
		} else {
			return videoResult, err
		}
	}
	return videoResult, nil
}

func (v *videoResultRepository) FindAll() ([]model.VideoResult, error) {
	var videoResult []model.VideoResult
	result := v.db.Find(&videoResult)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return videoResult, nil
		} else {
			return videoResult, err
		}
	}
	return videoResult, nil
}

func (v *videoResultRepository) UpdateByID(videoResult *model.VideoResult, by map[string]interface{}) error {
	result := v.db.Model(videoResult).Updates(by).Error
	return result
}

func (v *videoResultRepository) Delete(videoResult *model.VideoResult) error {
	result := v.db.Delete(videoResult).Error
	return result
}

func NewVideoResultRepository(db *gorm.DB) VideoResultRepository {
	repo := new(videoResultRepository)
	repo.db = db
	return repo
}
