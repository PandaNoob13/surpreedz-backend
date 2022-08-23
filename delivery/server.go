package delivery

import (
	"surpreedz-backend/config"
	"surpreedz-backend/delivery/controller"
	"surpreedz-backend/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	infra          manager.Infra
	managerUseCase manager.UseCaseManager
	engine         *gin.Engine
	host           string
	// tokenService utils.Token
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	managerRepo := manager.NewRepositoryManager(infra)
	managerUseCase := manager.NewUseCaseManager(managerRepo)
	host := appConfig.Url
	//tokenService := utils.NewTokenService(appConfig.TokenConfig)
	return &appServer{
		infra:          infra,
		managerUseCase: managerUseCase,
		engine:         r,
		host:           host,
		//tokenService: tokenService,
	}
}

func (a *appServer) initControllers() {
	controller.NewServiceController(a.engine, a.managerUseCase.AddService(), a.managerUseCase.FindService(), a.managerUseCase.UpdateService())
	controller.NewOrderController(a.engine, a.managerUseCase.AddOrder())
	controller.NewOrderStatusController(a.engine, a.managerUseCase.AddOrderStatus())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
