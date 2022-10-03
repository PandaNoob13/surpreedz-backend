package controller

import (
	"net/http"
	"surpreedz-backend/usecase"

	"github.com/gin-gonic/gin"
)

type SendAccCMSController struct {
	router                 *gin.Engine
	ucFindAccWithCondition usecase.FindAccountCMSUseCase
}

func (s *SendAccCMSController) GetVerifiedAcc(ctx *gin.Context) {
	//result, _, err := v.ucFindAccountByEmail.FindAccountByEmail(input.Email)

	result, url, err := s.ucFindAccWithCondition.FindAccountWithCondition("mst_account_detail.verified_status", "true")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "FAILED",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   "SUCCESS",
		"accounts": result,
		"photos":   url,
	})
}

func (s *SendAccCMSController) GetUnverifiedAcc(ctx *gin.Context) {
	result, url, err := s.ucFindAccWithCondition.FindAccountWithCondition("mst_account_detail.verified_status", "false")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "FAILED",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   "SUCCESS",
		"accounts": result,
		"photos":   url,
	})
}

func (s *SendAccCMSController) GetVerifiedReqAcc(ctx *gin.Context) {
	result, url, err := s.ucFindAccWithCondition.FindAccountWithCondition("mst_account_detail.verified_request", "waiting")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "FAILED",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   "SUCCESS",
		"accounts": result,
		"photos":   url,
	})
}

func NewSendAccCMSController(router *gin.Engine, ucFindAccWithCondition usecase.FindAccountCMSUseCase) *SendAccCMSController {

	controller := SendAccCMSController{
		router: router,
		ucFindAccWithCondition: ucFindAccWithCondition,
	}

	rCMSAcc := router.Group("/cms")
	{
		rCMSAcc.GET("/account-verified", controller.GetVerifiedAcc)
		rCMSAcc.GET("/account-unverified", controller.GetUnverifiedAcc)
		rCMSAcc.GET("/account-req-verified", controller.GetVerifiedReqAcc)
	}
	return &controller

}

