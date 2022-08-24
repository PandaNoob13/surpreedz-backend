package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type RetrieveAllFeedbackUseCase interface {
	RetrieveAllFeedback(page int, itemPerPage int) ([]model.Feedback, error)
}

type retrieveAllFeedbackUseCase struct {
	feedbackRepo repository.FeedbackRepository
}

func (r *retrieveAllFeedbackUseCase) RetrieveAllFeedback(page int, itemPerPage int) ([]model.Feedback, error) {
	return r.feedbackRepo.FindAll(page, itemPerPage)
}

func NewRetrieveAllFeedbackUseCase(feedbackRepo repository.FeedbackRepository) RetrieveAllFeedbackUseCase {
	return &retrieveAllFeedbackUseCase{
		feedbackRepo: feedbackRepo,
	}
}
