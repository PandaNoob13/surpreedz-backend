package controller

import (
	"strconv"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/model"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type VideoResultController struct {
	router         *gin.Engine
	insVidResUc    usecase.AddVideoResultUseCase
	rtAllVidResUc  usecase.RetrieveAllVideoResultUseCase
	fdVidResByIdUc usecase.FindVideoResultByIdUseCase
	api.BaseApi
}

func (v *VideoResultController) InsertVideoResult(c *gin.Context) {
	var addVideoResult model.VideoResult
	err := v.ParseRequestBody(c, &addVideoResult)
	if err != nil {
		v.Failed(c, utils.RequiredError())
		return
	}
	err = v.insVidResUc.AddVideoResult(&addVideoResult)
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

func (v *VideoResultController) FindVideoResultById(c *gin.Context) {
	videoResultId := c.Param("videoResultId")
	vidResId, _ := strconv.Atoi(videoResultId)
	vr, err := v.fdVidResByIdUc.FindVideoResultById(vidResId)
	if err != nil {
		v.Failed(c, err)
		return
	}
	v.Success(c, vr)
}

func NewVideoResultController(router *gin.Engine, insVidResUc usecase.AddVideoResultUseCase, rtAllVidResUc usecase.RetrieveAllVideoResultUseCase, fdVidResByIdUc usecase.FindVideoResultByIdUseCase) *VideoResultController {
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
		routerVideoResult.GET("get-video-result-by-id/:videoResultId", controller.FindVideoResultById)
	}
	return &controller
}
