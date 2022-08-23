package usecase

import (
	"surpreedz-backend/model"
	"surpreedz-backend/repository"
)

type AddVideoResultUseCase interface {
	AddVideoResult(videoResult *model.VideoResult) error
}

type addVideoResultUseCase struct {
	videoResultRepo repository.VideoResultRepository
}

func (v *addVideoResultUseCase) AddVideoResult(videoResult *model.VideoResult) error {
	return v.videoResultRepo.Create(videoResult)
}

func NewAddVideoResultUseCase(videoResultRepo repository.VideoResultRepository) AddVideoResultUseCase {
	return &addVideoResultUseCase{
		videoResultRepo: videoResultRepo,
	}
}
