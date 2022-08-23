package controller

import (
	"strconv"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/dto"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type OrderStatusController struct {
	router       *gin.Engine
	insOrdStatUc usecase.InsertOrderStatusUseCase
	api.BaseApi
}

func (os *OrderStatusController) InsertOrderStatus(c *gin.Context) {
	var addOrderStatus dto.OrderStatusDto
	orderStatusId := c.Param("orderStatusId")
	ordStatId, _ := strconv.Atoi(orderStatusId)
	err := os.ParseRequestBody(c, &addOrderStatus)
	if err != nil {
		os.Failed(c, utils.RequiredError())
		return
	}
	err = os.insOrdStatUc.AddOrderStatus(ordStatId, addOrderStatus.OrderId, addOrderStatus.Status, addOrderStatus.ResonOfRefund)
	if err != nil {
		os.Failed(c, err)
		return
	}
	os.Success(c, addOrderStatus)
}

func NewOrderStatusController(router *gin.Engine, insOrdStatUc usecase.InsertOrderStatusUseCase) *OrderStatusController {
	controller := OrderStatusController{
		router:       router,
		insOrdStatUc: insOrdStatUc,
	}
	routerOrderStatus := router.Group("/order-status")
	{
		routerOrderStatus.POST("/create-order-status/:orderStatusId", controller.InsertOrderStatus)
	}
	return &controller
}
