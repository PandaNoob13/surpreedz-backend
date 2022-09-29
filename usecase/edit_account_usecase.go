package usecase

import (
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
)

type EditAccountUsecase interface {
	EditProfile(editProfileDto *dto.EditProfileDto) error
	EditPassword(EditPasswordDto *dto.EditPasswordDto) error
	EditVerifiedStatus(VerifyFromCMSDto *dto.VerifyFromCMS) error
}

type editAccountUsecase struct {
	editAccountRepo repository.EditAccountRepository
}

func (e *editAccountUsecase) EditProfile(editProfile *dto.EditProfileDto) error {
	return e.editAccountRepo.EditProfile(editProfile)
}

func (e *editAccountUsecase) EditPassword(editPassword *dto.EditPasswordDto) error {
	return e.editAccountRepo.EditPassword(editPassword)
}

func (e *editAccountUsecase) EditVerifiedStatus(VerifyFromCMSDto *dto.VerifyFromCMS) error {
	return e.editAccountRepo.EditVerifiedStatus(VerifyFromCMSDto)
}

func NewEditAccountUsecase(editAccountRepo repository.EditAccountRepository) EditAccountUsecase {
	return &editAccountUsecase{
		editAccountRepo: editAccountRepo,
	}
}
