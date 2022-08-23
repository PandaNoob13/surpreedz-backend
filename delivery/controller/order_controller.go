package controller

import (
	"strconv"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/dto"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	router   *gin.Engine
	insOrdUc usecase.InsertOrderUseCase
	api.BaseApi
}

func (o *OrderController) InsertOrder(c *gin.Context) {
	var addOrder dto.OrderDto
	orderId := c.Param("orderId")
	ordId, _ := strconv.Atoi(orderId)
	err := o.ParseRequestBody(c, &addOrder)
	if err != nil {
		o.Failed(c, utils.RequiredError())
		return
	}
	err = o.insOrdUc.AddOrder(ordId, addOrder.BuyerId, addOrder.ServiceDetailId, addOrder.DueDate, addOrder.Occassion, addOrder.RecipientName, addOrder.Message, addOrder.RecipientDescription)
	if err != nil {
		o.Failed(c, err)
		return
	}
	o.Success(c, addOrder)
}

func NewOrderController(router *gin.Engine, insOrdUc usecase.InsertOrderUseCase) *OrderController {
	controller := OrderController{
		router:   router,
		insOrdUc: insOrdUc,
	}
	routerOrder := router.Group("/order")
	{
		routerOrder.POST("/create-order/:orderId", controller.InsertOrder)
	}
	return &controller
}
