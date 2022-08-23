package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindFeedbackByIdUseCase interface {
	FindFeedbackById(id int) (model.Feedback, error)
}

type findFeedbackByIdUseCase struct {
	feedbackRepo repository.FeedbackRepository
}

func (f *findFeedbackByIdUseCase) FindFeedbackById(id int) (model.Feedback, error) {
	return f.feedbackRepo.FindById(id)
}

func NewFindFeedbackByIdUseCase(feedbackRepo repository.FeedbackRepository) FindFeedbackByIdUseCase {
	return &findFeedbackByIdUseCase{
		feedbackRepo: feedbackRepo,
	}
}
