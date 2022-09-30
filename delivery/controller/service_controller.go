package controller

import (
	"net/http"
	"strconv"
	"surpreedz-backend/delivery/api"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	router    *gin.Engine
	insServUc usecase.InsertServiceUseCase
	fdSerUc   usecase.FindServiceUseCase
	hPgSrUc   usecase.ShowServicesHomePageUseCase
	updSerUc  usecase.UpdateServiceUseCase
	api.BaseApi
}

func (s *ServiceController) InsertService(c *gin.Context) {
	var addService dto.ServiceDto
	err := s.ParseRequestBody(c, &addService)
	if err != nil {
		s.Failed(c, utils.RequiredError())
		return
	}
	serviceDetailId, err := s.insServUc.AddService(&addService)
	if err != nil {
		s.Failed(c, err)
		return
	}
	var serviceDetailIdDto dto.CreateServiceResponseDto
	serviceDetailIdDto.Service_Detail_Id = serviceDetailId
	s.Success(c, serviceDetailIdDto)
}

func (s *ServiceController) RetrieveHomePage(c *gin.Context) {
	page := c.Query("page")
	pg, _ := strconv.Atoi(page)
	limit := c.Query("limit")
	lm, _ := strconv.Atoi(limit)
	services, err := s.hPgSrUc.HomePageRetrieveAll(pg, lm)
	if err != nil {
		s.Failed(c, err)
		return
	}
	s.Success(c, services)
}

func (s *ServiceController) UpdateService(c *gin.Context) {
	serviceId := c.Query("serviceId")
	servId, _ := strconv.Atoi(serviceId)
	sh, err := s.fdSerUc.FindServiceById(servId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		// s.Failed(c, err)
		return
	} else {
		// s.Success(c, sh)
		var existService dto.EditServiceDto
		err := s.ParseRequestBody(c, &existService)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAILED",
				"message": err.Error(),
			})
			// s.Failed(c, utils.RequiredError())
			return
		}
		err = s.updSerUc.EditService(sh.ID, &existService)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": err.Error(),
			})
			// s.Failed(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"message": "Success updating service",
		})
		// s.Success(c, existService)
	}
}

func (s *ServiceController) FindServiceById(c *gin.Context) {
	serviceId := c.Param("serviceId")
	servId, _ := strconv.Atoi(serviceId)
	serv, err := s.fdSerUc.FindServiceById(servId)
	if err != nil {
		s.Failed(c, err)
		return
	}
	s.Success(c, serv)
}

func NewServiceController(router *gin.Engine, insSerUc usecase.InsertServiceUseCase, fdSerUc usecase.FindServiceUseCase, updSerUc usecase.UpdateServiceUseCase, hPGSrUc usecase.ShowServicesHomePageUseCase) *ServiceController {
	contoller := ServiceController{
		router:    router,
		insServUc: insSerUc,
		fdSerUc:   fdSerUc,
		updSerUc:  updSerUc,
		hPgSrUc:   hPGSrUc,
	}
	routerService := router.Group("/service-detail")
	{
		routerService.POST("/create-service-detail", contoller.InsertService)
		routerService.PUT("/edit-service-detail/", contoller.UpdateService)
		routerService.GET("/homepage/", contoller.RetrieveHomePage)
		routerService.GET("/get-service-detail/:serviceId", contoller.FindServiceById)
	}
	return &contoller
}
