package delivery

import (
	"surpreedz-backend/config"
	"surpreedz-backend/delivery/controller"
	"surpreedz-backend/manager"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	managerRepo    manager.RepositoryManager
	infra          manager.Infra
	managerUsecase manager.UseCaseManager
	engine         *gin.Engine
	tokenService   utils.Token
	host           string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	managerRepo := manager.NewRepositoryManager(infra)
	managerUsecase := manager.NewUseCaseManager(managerRepo)
	host := appConfig.Url
	tokenService := utils.NewTokenService(appConfig.TokenConfig)
	return &appServer{
		managerRepo:    managerRepo,
		infra:          infra,
		managerUsecase: managerUsecase,
		engine:         r,
		host:           host,
		tokenService:   tokenService,
	}
}

func (a *appServer) initControllers() {
	controller.NewServiceController(a.engine, a.managerUsecase.AddService(), a.managerUsecase.FindService(), a.managerUsecase.UpdateService(), a.managerUsecase.RetrieveServiceHomePage())
	controller.NewOrderController(a.engine, a.managerUsecase.AddOrder(), a.managerUsecase.RetrieveAllOrder(), a.managerUsecase.FindOrderById())
	controller.NewOrderStatusController(a.engine, a.managerUsecase.AddOrderStatus())
	controller.NewVideoResultController(a.engine, a.managerUsecase.AddVideoResult(), a.managerUsecase.RetrieveAllVideoResult(), a.managerUsecase.FindVideoResultById())
	controller.NewFeedbackController(a.engine, a.managerUsecase.AddFeedback(), a.managerUsecase.RetrieveAllFeedback(), a.managerUsecase.FindFeedbackById())
	controller.NewLoginController(a.engine, a.tokenService, a.managerUsecase.FindAccountUseCase(), a.managerUsecase.FindPasswordByAccId())
	controller.NewSignUpController(a.engine, a.managerUsecase.SignUpAccountUseCase(), a.managerUsecase.FindAccountUseCase())
	controller.NewEditAccountController(a.engine, a.tokenService, a.managerUsecase.EditAccountInfoUsecase(), a.managerUsecase.FindPasswordByAccId())
	controller.NewPaymentController(a.engine)
	controller.NewPaymentStatusController(a.engine, a.managerUsecase.AddPaymentStatus())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
