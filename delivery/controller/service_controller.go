package controller

import (
	"net/http"
	"strconv"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/dto"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	router    *gin.Engine
	insServUc usecase.InsertServiceUseCase
	fdSerUc   usecase.FindServiceUseCase
	updSerUc  usecase.UpdateServiceUseCase
	api.BaseApi
}

func (s *ServiceController) InsertService(c *gin.Context) {
	var addService dto.ServiceDto
	serviceId := c.Param("serviceId")
	servId, _ := strconv.Atoi(serviceId)
	err := s.ParseRequestBody(c, &addService)
	if err != nil {
		s.Failed(c, utils.RequiredError())
		return
	}
	err = s.insServUc.AddService(servId, addService.SellerId, addService.Role, addService.Description, addService.Price, addService.VideoLink)
	if err != nil {
		s.Failed(c, err)
		return
	}
	s.Success(c, addService)
}

func (s *ServiceController) UpdateService(c *gin.Context) {
	serviceId := c.Param("serviceId")
	servId, _ := strconv.Atoi(serviceId)
	sh, err := s.fdSerUc.FindServiceById(servId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		var by map[string]interface{}
		if err := c.ShouldBindJSON(&by); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAILED",
				"message": err.Error(),
			})
		} else {
			err := s.updSerUc.EditService(&sh, by)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"status":  "FAILED",
					"message": "error when updating service",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  "SUCCESS",
					"message": sh,
				})
			}
		}
	}
}

func NewServiceController(router *gin.Engine, insSerUc usecase.InsertServiceUseCase, fdSerUc usecase.FindServiceUseCase, updSerUc usecase.UpdateServiceUseCase) *ServiceController {
	contoller := ServiceController{
		router:    router,
		insServUc: insSerUc,
		fdSerUc:   fdSerUc,
		updSerUc:  updSerUc,
	}
	routerService := router.Group("/service-detail")
	{
		routerService.POST("/create-service-detail/:serviceId", contoller.InsertService)
		routerService.PUT("/edit-service-detail/:serviceId", contoller.UpdateService)
	}
	return &contoller
}
