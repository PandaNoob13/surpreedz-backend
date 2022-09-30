package controller

import (
	"strconv"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type PaymentStatusController struct {
	router       *gin.Engine
	insPayStatUc usecase.AddPaymentStatusUseCase
	api.BaseApi
}

func (ps *PaymentStatusController) AddPaymentStatus(c *gin.Context) {
	var addPaymentStatus dto.PaymentStatusDto
	err := ps.ParseRequestBody(c, &addPaymentStatus)
	if err != nil {
		ps.Failed(c, utils.RequiredError())
		return
	}
	stringInt, _ := strconv.Atoi(addPaymentStatus.OrderId)
	paymentStatus := model.PaymentStatus{
		OrderId:       stringInt,
		PaymentType:   addPaymentStatus.PaymentType,
		StatusPayment: addPaymentStatus.StatusPayment,
	}
	paymentStatus.TimeUpdated = time.Now()
	err = ps.insPayStatUc.AddPaymentStatus(&paymentStatus)
	if err != nil {
		ps.Failed(c, err)
		return
	}
	ps.Success(c, paymentStatus)
}

func NewPaymentStatusController(router *gin.Engine, insPayStatUc usecase.AddPaymentStatusUseCase) *PaymentStatusController {
	controller := PaymentStatusController{
		router:       router,
		insPayStatUc: insPayStatUc,
	}
	routerPaymentStatus := router.Group("/order")
	{
		routerPaymentStatus.POST("/update-payment-status", controller.AddPaymentStatus)
	}
	return &controller
}
