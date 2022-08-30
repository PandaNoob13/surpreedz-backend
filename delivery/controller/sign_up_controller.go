package controller

import (
	"net/http"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	router               *gin.Engine
	ucSignUpAccount      usecase.SignUpUsecase
	ucFindAccountByEmail usecase.FindAccountUseCase
}

func (s *SignUpController) buatAkunBaru(ctx *gin.Context) {

	var input dto.AccountFormInfo
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "FAILED",
			"message": "can't bind struct",
		})
		return
	}

	_, _, err := s.ucFindAccountByEmail.FindAccountByEmail(input.Email)
	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Account exist with this email",
		})
		return
	}

	err = s.ucSignUpAccount.SignUpNewAccount(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}

	AccRes, _, err := s.ucFindAccountByEmail.FindAccountByEmail(input.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"account": AccRes,
	})
}

func NewSignUpController(router *gin.Engine, ucSignUpAccount usecase.SignUpUsecase, ucFindAccountByEmail usecase.FindAccountUseCase) *SignUpController {

	controller := SignUpController{
		router:               router,
		ucSignUpAccount:      ucSignUpAccount,
		ucFindAccountByEmail: ucFindAccountByEmail,
	}

	rSignUp := router.Group("/account")
	{
		rSignUp.POST("/signup", controller.buatAkunBaru)
	}
	return &controller

}
