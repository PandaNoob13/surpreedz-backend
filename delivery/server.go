package delivery

import(
		"surpreedz-backend/config"
		"surpreedz-backend/manager"

		"github.com/gin-gonic/gin"
)

type appServer struct {
	// managerUscs manager.Usecase
	engine *gin.Engine
	infra manager.Infra
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
		//managerUscs:  managerUseCase,
		engine:       r,
		host:         host,
		infra: infra,
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