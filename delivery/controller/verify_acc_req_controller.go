package controller

import (
	"net/http"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"

	"github.com/gin-gonic/gin"
)

type VerifyAccController struct {
	router *gin.Engine
	//ucFindAccountByEmail usecase.FindAccountUseCase
	ucEditAcc usecase.EditAccountUsecase
	api.BaseApi
}

func (v *VerifyAccController) VerifikasiAkun(ctx *gin.Context) {
	var accIdInput dto.AccIdInput
	// if err := v.ParseRequestBody(ctx, &accIdInput); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  "FAILED",
	// 		"message": "can't bind struct",
	// 	})
	// 	return
	// }

	if err := ctx.BindJSON(&accIdInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "FAILED",
			"message": "can't bind struct",
		})
		return
	}

	var data dto.VerifyFromCMS
	data.AccountId = accIdInput.AccountId
	data.VerifiedRequest = "waiting"
	data.VerifiedStatus = false

	//result, _, err := v.ucFindAccountByEmail.FindAccountByEmail(input.Email)

	err := v.ucEditAcc.EditVerifiedStatus(&data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Account not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		// "account_id": result.ID,
		// "photo":      result.AccountDetail.PhotoProfiles[len(result.AccountDetail.PhotoProfiles)-1].PhotoLink,
	})
}

func NewVerifyAccController(router *gin.Engine /* ucFindAccountByEmail usecase.FindAccountUseCase, */, ucEditAcc usecase.EditAccountUsecase) *VerifyAccController {

	controller := VerifyAccController{
		router: router,
		//ucFindAccountByEmail: ucFindAccountByEmail,
		ucEditAcc: ucEditAcc,
	}

	rVerifyAcc := router.Group("/verify")
	{
		rVerifyAcc.PUT("/account", controller.VerifikasiAkun)
	}
	return &controller

}
