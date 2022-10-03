package controller

import (
	"net/http"
	"surpreedz-backend/usecase"

	"github.com/gin-gonic/gin"
)

type SendBuyerSellerAccCMSController struct {
	router               *gin.Engine
	ucFindBuyerSellerAcc usecase.FindBuyerSellerAccCMSUseCase
}

func (s *SendBuyerSellerAccCMSController) GetBuyerAcc(ctx *gin.Context) {
	result, url, err := s.ucFindBuyerSellerAcc.FindBuyerAcc()
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

func (s *SendBuyerSellerAccCMSController) GetSellererAcc(ctx *gin.Context) {
	result, url, err := s.ucFindBuyerSellerAcc.FindSellerAcc()
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

func (s *SendBuyerSellerAccCMSController) GetBuyerSellerAcc(ctx *gin.Context) {
	result, url, err := s.ucFindBuyerSellerAcc.FindBuyerSellerAcc()
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

func NewSendBuyerSellerAccCMSController(router *gin.Engine, ucFindBuyerSellerAcc usecase.FindBuyerSellerAccCMSUseCase) *SendBuyerSellerAccCMSController {

	controller := SendBuyerSellerAccCMSController{
		router:               router,
		ucFindBuyerSellerAcc: ucFindBuyerSellerAcc,
	}

	rCMSAcc := router.Group("/cms")
	{
		rCMSAcc.GET("/buyer-account", controller.GetBuyerAcc)
		rCMSAcc.GET("/seller-account", controller.GetSellererAcc)
		rCMSAcc.GET("/buyer-seller-account", controller.GetBuyerSellerAcc)
	}
	return &controller

}
