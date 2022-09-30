package controller

import (
	"strconv"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type VideoResultController struct {
	router         *gin.Engine
	insVidResUc    usecase.AddVideoResultUseCase
	rtAllVidResUc  usecase.RetrieveAllVideoResultUseCase
	fdVidResByIdUc usecase.FindVideoResultByOrderIdUseCase
	api.BaseApi
}

func (v *VideoResultController) InsertVideoResult(c *gin.Context) {
	var addVideoResult dto.VideoResultDto
	err := v.ParseRequestBody(c, &addVideoResult)
	if err != nil {
		v.Failed(c, utils.RequiredError())
		return
	}
	videoResult := model.VideoResult{
		OrderId:   addVideoResult.OrderId,
		VideoLink: addVideoResult.VideoLink,
	}
	err = v.insVidResUc.AddVideoResult(&videoResult, addVideoResult.DataUrl)
	if err != nil {
		v.Failed(c, err)
		return
	}
	v.Success(c, addVideoResult)
}

func (v *VideoResultController) RetriveAllVideoResult(c *gin.Context) {
	page := c.Param("page")
	pg, _ := strconv.Atoi(page)
	limit := c.Param("limit")
	lm, _ := strconv.Atoi(limit)
	videoResults, err := v.rtAllVidResUc.RetriveAllVideoResult(pg, lm)
	if err != nil {
		v.Failed(c, err)
		return
	}
	v.Success(c, videoResults)
}

func (v *VideoResultController) FindVideoResultByOrderId(c *gin.Context) {
	videoResultId := c.Query("orderId")
	vr, err := v.fdVidResByIdUc.FindVideoResultByOrderId(videoResultId)
	if err != nil {
		v.Failed(c, err)
		return
	}
	v.Success(c, vr)
}

func NewVideoResultController(router *gin.Engine, insVidResUc usecase.AddVideoResultUseCase, rtAllVidResUc usecase.RetrieveAllVideoResultUseCase, fdVidResByIdUc usecase.FindVideoResultByOrderIdUseCase) *VideoResultController {
	controller := VideoResultController{
		router:         router,
		insVidResUc:    insVidResUc,
		rtAllVidResUc:  rtAllVidResUc,
		fdVidResByIdUc: fdVidResByIdUc,
	}
	routerVideoResult := router.Group("/video-result")
	{
		routerVideoResult.POST("/create-video-result", controller.InsertVideoResult)
		routerVideoResult.GET("/get-video-result/:page/:limit", controller.RetriveAllVideoResult)
		routerVideoResult.GET("/get-video-result-by-order-id", controller.FindVideoResultByOrderId)
	}
	return &controller
}
