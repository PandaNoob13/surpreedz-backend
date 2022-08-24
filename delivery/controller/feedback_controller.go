package controller

import (
	"strconv"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/model"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type FeedbackController struct {
	router     *gin.Engine
	insFdBcUc  usecase.AddFeedbackUseCase
	fdFdBcUc   usecase.FindFeedbackByIdUseCase
	reAlFdBcUc usecase.RetrieveAllFeedbackUseCase
	api.BaseApi
}

func (f *FeedbackController) InsertFeedback(c *gin.Context) {
	var addFeedback model.Feedback
	err := f.ParseRequestBody(c, &addFeedback)
	if err != nil {
		f.Failed(c, utils.RequiredError())
		return
	}
	err = f.insFdBcUc.AddFeedback(&addFeedback)
	if err != nil {
		f.Failed(c, err)
		return
	}
	f.Success(c, addFeedback)
}

func (f *FeedbackController) RetrieveAllFeedback(c *gin.Context) {
	page := c.Param("page")
	pg, _ := strconv.Atoi(page)
	limit := c.Param("limit")
	lm, _ := strconv.Atoi(limit)
	feedbacks, err := f.reAlFdBcUc.RetrieveAllFeedback(pg, lm)
	if err != nil {
		f.Failed(c, err)
		return
	}
	f.Success(c, feedbacks)
}

func (f *FeedbackController) FindFeedbackById(c *gin.Context) {
	feedbackId := c.Param("feedbackId")
	fbId, _ := strconv.Atoi(feedbackId)
	feedback, err := f.fdFdBcUc.FindFeedbackById(fbId)
	if err != nil {
		f.Failed(c, err)
		return
	}
	f.Success(c, feedback)
}

func NewFeedbackController(router *gin.Engine, insFdBcUc usecase.AddFeedbackUseCase, reAlFdBcUc usecase.RetrieveAllFeedbackUseCase, fdFdBcUc usecase.FindFeedbackByIdUseCase) *FeedbackController {
	controller := FeedbackController{
		router:     router,
		insFdBcUc:  insFdBcUc,
		reAlFdBcUc: reAlFdBcUc,
		fdFdBcUc:   fdFdBcUc,
	}
	routerFeedback := router.Group("/feedback")
	{
		routerFeedback.POST("/create-feedback", controller.InsertFeedback)
		routerFeedback.GET("/get-feedback/:page/:limit", controller.RetrieveAllFeedback)
		routerFeedback.GET("get-feedback-by-id/:feedbackId", controller.FindFeedbackById)
	}
	return &controller
}
