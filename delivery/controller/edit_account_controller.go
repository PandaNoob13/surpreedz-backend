package controller

import (
	"fmt"
	"net/http"
	"surpreedz-backend/delivery/middleware"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type EditAccountController struct {
	router            *gin.Engine
	tokenService      utils.Token
	ucEditAcc         usecase.EditAccountUsecase
	ucFindPassByAccid usecase.FindPasswordUseCase
}

func (e *EditAccountController) EditAccountPassword(ctx *gin.Context) {
	var input dto.EditPasswordDto
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	PassRes, err := e.ucFindPassByAccid.FindPasswordByAccountId(input.AccountId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}

	if PassRes.Password != input.OldPassword {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "FAILED",
			"message": "old password did not match",
		})
		return
	}

	err = e.ucEditAcc.EditPassword(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
	})
}

func (e *EditAccountController) EditAccountProfile(ctx *gin.Context) {
	var input dto.EditProfileDto
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	} else {
		fmt.Println("IF YOU CAN'T BIND IT, WHY DON'T YOU TELL ME?!?!?!?")
		fmt.Println(input)
	}

	err := e.ucEditAcc.EditProfile(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
	})
}

func NewEditAccountController(router *gin.Engine, tokenService utils.Token, ucEditAcc usecase.EditAccountUsecase, ucFindPassByAccid usecase.FindPasswordUseCase) *EditAccountController {

	controller := EditAccountController{
		router:            router,
		tokenService:      tokenService,
		ucEditAcc:         ucEditAcc,
		ucFindPassByAccid: ucFindPassByAccid,
	}

	rEditAcc := router.Group("/account", middleware.NewTokenValidator(tokenService).RequireToken())
	{
		rEditAcc.PUT("/edit-password", controller.EditAccountPassword)
		rEditAcc.PUT("/edit-profile", controller.EditAccountProfile)
	}
	return &controller

}
