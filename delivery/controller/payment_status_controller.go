package controller

import (
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/model"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type PaymentStatusController struct {
	router       *gin.Engine
	insPayStatUc usecase.AddPaymentStatusUseCase
	api.BaseApi
}

func (ps *PaymentStatusController) AddPaymentStatus(c *gin.Context) {
	var addPaymentStatus model.PaymentStatus
	err := ps.ParseRequestBody(c, &addPaymentStatus)
	if err != nil {
		ps.Failed(c, utils.RequiredError())
		return
	}
	err = ps.insPayStatUc.AddPaymentStatus(&addPaymentStatus)
	if err != nil {
		ps.Failed(c, err)
		return
	}
	ps.Success(c, addPaymentStatus)
}

func NewPaymentStatusController(router *gin.Engine, insPayStatUc usecase.AddPaymentStatusUseCase) *PaymentStatusController {
	controller := PaymentStatusController{
		router:       router,
		insPayStatUc: insPayStatUc,
	}
	routerPaymentStatus := router.Group("/order")
	{
		routerPaymentStatus.POST("/payment-status", controller.AddPaymentStatus)
	}
	return &controller
}
