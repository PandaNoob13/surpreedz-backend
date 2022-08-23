package delivery

import (
	"surpreedz-backend/config"
	"surpreedz-backend/delivery/controller"
	"surpreedz-backend/manager"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	managerRepo  manager.RepositoryManager
	infra        manager.Infra
	managerUsecase  manager.UseCaseManager
	engine       *gin.Engine
	tokenService utils.Token
	host         string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	managerRepo := manager.NewRepositoryManager(infra)
	managerUseCase := manager.NewUseCaseManager(managerRepo)
	host := appConfig.Url
	tokenService := utils.NewTokenService(appConfig.TokenConfig)
	return &appServer{
		managerRepo:  managerRepo,
		infra:        infra,
		managerUsecase:  managerUseCase,
		engine:       r,
		host:         host,
		tokenService: tokenService,
	}
}

func (a *appServer) initControllers() {
	controller.NewServiceController(a.engine, a.managerUseCase.AddService(), a.managerUseCase.FindService(), a.managerUseCase.UpdateService())
	controller.NewOrderController(a.engine, a.managerUseCase.AddOrder())
	controller.NewOrderStatusController(a.engine, a.managerUseCase.AddOrderStatus())
	controller.NewLoginController(a.engine, a.tokenService, a.managerUseCase.FindAccountUseCase())
	controller.NewSignUpController(a.engine, a.managerUscs.SignUpAccountUseCase(), a.managerUseCase.FindAccountUseCase())
	controller.NewEditAccountController(a.engine, a.tokenService, a.managerUseCase.EditAccountInfoUsecase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
