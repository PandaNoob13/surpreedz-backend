package controller

import (
	"fmt"
	"strconv"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"

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
	fmt.Println(addOrder)
	if err != nil {
		o.Failed(c, err)
		return
	}
	order_id, err := o.insOrdUc.AddOrder(addOrder)
	if err != nil {
		o.Failed(c, err)
		return
	}
	var response dto.CreateOrderResponseDto
	response.Order_id = order_id
	o.Success(c, response)
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

func (o *OrderController) FindAllOrderByBuyerId(c *gin.Context) {
	buyerId := c.Query("buyerId")
	byrId, _ := strconv.Atoi(buyerId)
	order, err := o.fdOrdByIdUc.FindAllOrderByBuyerId(byrId)
	if err != nil {
		o.Failed(c, err)
		return
	}
	o.Success(c, order)
}

func (o *OrderController) FindAllOrderByServiceDetailId(c *gin.Context) {
	buyerId := c.Query("serviceDetailId")
	byrId, _ := strconv.Atoi(buyerId)
	order, err := o.fdOrdByIdUc.FindAllOrderByServiceDetailId(byrId)
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
		routerOrder.GET("/get-order/:page/:limit", controller.RetrieveAllOrder)
		routerOrder.GET("/get-order-by-id/:orderId", controller.FindOrderById)
		routerOrder.GET("/get-all-order-by-buyer-id", controller.FindAllOrderByBuyerId)
		routerOrder.GET("/get-all-order-by-service-detail-id", controller.FindAllOrderByServiceDetailId)

	}
	return &controller
}
