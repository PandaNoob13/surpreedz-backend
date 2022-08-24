package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type RetrieveAllVideoResultUseCase interface {
	RetriveAllVideoResult(page int, itemPerPage int) ([]model.VideoResult, error)
}

type retrieveAllVideoResultUseCase struct {
	videoResultRepo repository.VideoResultRepository
}

func (r *retrieveAllVideoResultUseCase) RetriveAllVideoResult(page int, itemPerPage int) ([]model.VideoResult, error) {
	return r.videoResultRepo.FindAll(page, itemPerPage)
}

func NewRetrieveAllVideoResult(videoResultRepo repository.VideoResultRepository) RetrieveAllVideoResultUseCase {
	return &retrieveAllVideoResultUseCase{
		videoResultRepo: videoResultRepo,
	}
}
