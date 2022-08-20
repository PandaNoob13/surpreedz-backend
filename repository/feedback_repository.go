package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type FeedbackRepository interface {
	Create(feedback *model.Feedback) error
	FindById(id int) (model.Feedback, error)
	FindAll() ([]model.Feedback, error)
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
	result := f.db.First(&feedback, "order_id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return feedback, nil
		} else {
			return feedback, err
		}
	}
	return feedback, nil
}

func (f *feedbackRepository) FindAll() ([]model.Feedback, error) {
	var feedbacks []model.Feedback
	result := f.db.Find(&feedbacks)
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
