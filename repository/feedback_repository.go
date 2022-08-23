package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type FeedbackRepository interface {
	Create(feedback *model.Feedback) error
	FindById(id int) (model.Feedback, error)
	FindAll(page int, itemPerPage int) ([]model.Feedback, error)
	UpdateByID(feedback *model.Feedback, by map[string]interface{}) error
	Delete(feedback *model.Feedback) error
}

type feedbackRepository struct {
	db *gorm.DB
}

func (f *feedbackRepository) Create(feedback *model.Feedback) error {
	result := f.db.Create(feedback).Error
	return result
}

func (f *feedbackRepository) FindById(id int) (model.Feedback, error) {
	var feedback model.Feedback
	result := f.db.Where("mst_feedback.id = ?", id).First(&feedback)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return feedback, nil
		} else {
			return feedback, err
		}
	}
	return feedback, nil
}

func (f *feedbackRepository) FindAll(page int, itemPerPage int) ([]model.Feedback, error) {
	var feedbacks []model.Feedback
	offset := itemPerPage * (page - 1)
	result := f.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).First(&feedbacks)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return feedbacks, nil
		} else {
			return feedbacks, err
		}
	}
	return feedbacks, nil
}

func (f *feedbackRepository) UpdateByID(feedback *model.Feedback, by map[string]interface{}) error {
	result := f.db.Model(feedback).Updates(by).Error
	return result
}

func (f *feedbackRepository) Delete(feedback *model.Feedback) error {
	result := f.db.Delete(feedback).Error
	return result
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	repo := new(feedbackRepository)
	repo.db = db
	return repo
}
