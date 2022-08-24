package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type AddFeedbackUseCase interface {
	AddFeedback(feedback *model.Feedback) error
}

type addFeedbackUseCase struct {
	feedbackRepo repository.FeedbackRepository
}

func (a *addFeedbackUseCase) AddFeedback(feedback *model.Feedback) error {
	return a.feedbackRepo.Create(feedback)
}

func NewAddFeedbackUseCase(feedbackRepo repository.FeedbackRepository) AddFeedbackUseCase {
	return &addFeedbackUseCase{
		feedbackRepo: feedbackRepo,
	}
}
