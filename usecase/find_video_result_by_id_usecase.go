package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type FindVideoResultByIdUseCase interface {
	FindVideoResultById(id int) (model.VideoResult, error)
}

type findVideoResultByIdUseCase struct {
	videoResultRepo repository.VideoResultRepository
}

func (v *findVideoResultByIdUseCase) FindVideoResultById(id int) (model.VideoResult, error) {
	return v.videoResultRepo.FindById(id)
}

func NewFindVideoResultByIdUseCase(videoResultRepo repository.VideoResultRepository) FindVideoResultByIdUseCase {
	return &findVideoResultByIdUseCase{
		videoResultRepo: videoResultRepo,
	}
}
