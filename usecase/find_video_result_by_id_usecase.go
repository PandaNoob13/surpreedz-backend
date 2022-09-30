package usecase

import (
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
)

type FindVideoResultByOrderIdUseCase interface {
	FindVideoResultByOrderId(id string) (dto.VideoResultDto, error)
}

type findVideoResultByOrderIdUseCase struct {
	videoResultRepo repository.VideoResultRepository
}

func (v *findVideoResultByOrderIdUseCase) FindVideoResultByOrderId(id string) (dto.VideoResultDto, error) {
	return v.videoResultRepo.FindByOrderId(id)
}

func NewFindVideoResultByIdUseCase(videoResultRepo repository.VideoResultRepository) FindVideoResultByOrderIdUseCase {
	return &findVideoResultByOrderIdUseCase{
		videoResultRepo: videoResultRepo,
	}
}
