package delivery

import (
	"surpreedz-backend/config"
	"surpreedz-backend/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	infra manager.Infra
	// managerUscs manager.Usecase
	engine *gin.Engine
	// tokenService utils.Token
	host string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	//managerRepo := manager.NewRepositoryManager(infra)
	//managerUseCase := manager.NewUseCaseManager(managerRepo)
	host := appConfig.Url
	//tokenService := utils.NewTokenService(appConfig.TokenConfig)
	return &appServer{
		infra: infra,
		//managerUscs:  managerUseCase,
		engine: r,
		host:   host,
		//tokenService: tokenService,
	}
}

func (a *appServer) initControllers() {
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
