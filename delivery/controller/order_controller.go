package controller

import (
	"strconv"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	router      *gin.Engine
	insOrdUc    usecase.InsertOrderUseCase
	rtAllOrdUc  usecase.RetrieveAllOrderUseCase
	fdOrdByIdUc usecase.FindOrderByIdUseCase
	api.BaseApi
}

func (o *OrderController) InsertOrder(c *gin.Context) {
	var addOrder dto.OrderDto
	err := o.ParseRequestBody(c, &addOrder)
	if err != nil {
		o.Failed(c, utils.RequiredError())
		return
	}
	err = o.insOrdUc.AddOrder(addOrder)
	if err != nil {
		o.Failed(c, err)
		return
	}
	o.Success(c, addOrder)
}

func (o *OrderController) RetrieveAllOrder(c *gin.Context) {
	page := c.Param("page")
	pg, _ := strconv.Atoi(page)
	limit := c.Param("limit")
	lm, _ := strconv.Atoi(limit)
	orders, err := o.rtAllOrdUc.RetrieveAllOrder(pg, lm)
	if err != nil {
		o.Failed(c, err)
		return
	}
	o.Success(c, orders)
}

func (o *OrderController) FindOrderById(c *gin.Context) {
	orderId := c.Param("orderId")
	ordId, _ := strconv.Atoi(orderId)
	order, err := o.fdOrdByIdUc.FindOrderById(ordId)
	if err != nil {
		o.Failed(c, err)
		return
	}
	o.Success(c, order)
}

func NewOrderController(router *gin.Engine, insOrdUc usecase.InsertOrderUseCase, rtAllOrdUc usecase.RetrieveAllOrderUseCase, fdOrdByIdUc usecase.FindOrderByIdUseCase) *OrderController {
	controller := OrderController{
		router:      router,
		insOrdUc:    insOrdUc,
		rtAllOrdUc:  rtAllOrdUc,
		fdOrdByIdUc: fdOrdByIdUc,
	}
	routerOrder := router.Group("/order")
	{
		routerOrder.POST("/create-order", controller.InsertOrder)
		routerOrder.GET("get-order/:page/:limit", controller.RetrieveAllOrder)
		routerOrder.GET("get-order-by-id/:orderId", controller.FindOrderById)
	}
	return &controller
}
