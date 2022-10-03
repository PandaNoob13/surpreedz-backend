package controller

import (
	"net/http"
	"surpreedz-backend/usecase"

	"github.com/gin-gonic/gin"
)

type SendOrderCMSController struct {
	router         *gin.Engine
	ucFindOrderCMS usecase.FindOrderCMSUseCase
}

func (s *SendOrderCMSController) GetWaitingOrder(ctx *gin.Context) {
	//result, _, err := v.ucFindAccountByEmail.FindAccountByEmail(input.Email)

	result, err := s.ucFindOrderCMS.FindOrderWithCondition("mst_order_status.status", "Waiting for confirmation")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "FAILED",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"orders": result,
	})
}

func (s *SendOrderCMSController) GetOnProgressOrder(ctx *gin.Context) {
	result, err := s.ucFindOrderCMS.FindOrderWithCondition("mst_order_status.status", "On progress")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "FAILED",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"orders": result,
	})
}

func (s *SendOrderCMSController) GetSubmittedOrder(ctx *gin.Context) {
	result, err := s.ucFindOrderCMS.FindOrderWithCondition("mst_order_status.status", "Submitted")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "FAILED",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"orders": result,
	})
}

func NewSendOrderCMSController(router *gin.Engine, ucFindOrderCMS usecase.FindOrderCMSUseCase) *SendOrderCMSController {

	controller := SendOrderCMSController{
		router:         router,
		ucFindOrderCMS: ucFindOrderCMS,
	}

	rCMSAcc := router.Group("/cms")
	{
		rCMSAcc.GET("/order-waiting", controller.GetWaitingOrder)
		rCMSAcc.GET("/order-progress", controller.GetOnProgressOrder)
		rCMSAcc.GET("/order-submit", controller.GetSubmittedOrder)
	}
	return &controller

}
