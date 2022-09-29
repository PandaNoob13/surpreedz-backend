package controller

import (
	"net/http"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"

	"github.com/gin-gonic/gin"
)

type VerifyAccController struct {
	router               *gin.Engine
	ucFindAccountByEmail usecase.FindAccountUseCase
}

func (v *VerifyAccController) VerifikasiAkun(ctx *gin.Context) {
	var input dto.EmailInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "FAILED",
			"message": "can't bind struct",
		})
		return
	}

	result, _, err := v.ucFindAccountByEmail.FindAccountByEmail(input.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Account not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     "SUCCESS",
		"account_id": result.ID,
		"photo":      result.AccountDetail.PhotoProfiles[len(result.AccountDetail.PhotoProfiles)-1].PhotoLink,
	})
}

func NewVerifyAccController(router *gin.Engine, ucFindAccountByEmail usecase.FindAccountUseCase) *VerifyAccController {

	controller := VerifyAccController{
		router:               router,
		ucFindAccountByEmail: ucFindAccountByEmail,
	}

	rVerifyAcc := router.Group("/verify")
	{
		rVerifyAcc.POST("/account", controller.VerifikasiAkun)
	}
	return &controller

}
