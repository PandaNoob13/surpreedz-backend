package usecase

import (
	"surpreedz-backend/model/dto"
	"surpreedz-backend/repository"
)

type SignUpUsecase interface {
	SignUpNewAccount(accountFormInfo *dto.AccountFormInfo) error
}

type signUpUsecase struct {
	signUpRepo repository.SignUpRepository
}

func (s *signUpUsecase) SignUpNewAccount(accountFormInfo *dto.AccountFormInfo) error {
	return s.signUpRepo.SignUpAccount(accountFormInfo)
}

func NewSignUpUsecase(signUpRepo repository.SignUpRepository) SignUpUsecase {
	return &signUpUsecase{
		signUpRepo: signUpRepo,
	}
}
