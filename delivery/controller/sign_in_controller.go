package controller

import (
	"net/http"
	"surpreedz-backend/model"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	router           *gin.Engine
	tokenService     utils.Token
	ucFindAccByEmail usecase.FindAccountUseCase
}

func (l *LoginController) loginAkunCustomer(ctx *gin.Context) {

	var user model.Credential
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	AccRes, err := l.ucFindAccByEmail.FindAccountByEmail(user.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}

	if user.Username == AccRes.Email && user.Password == AccRes.Password {
		token, err := l.tokenService.CreateAccessToken(&user)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		err = l.tokenService.StoreAccessToken(user.Username, token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "email and password did not match",
		})
	}
}

func NewLoginController(router *gin.Engine, tokenService utils.Token, ucFindAccByEmail usecase.FindAccountUseCase) *LoginController {

	controller := LoginController{
		router:           router,
		tokenService:     tokenService,
		ucFindAccByEmail: ucFindAccByEmail,
	}

	rLogin := router.Group("/api")
	{
		rLogin.POST("/auth/login", controller.loginAkunCustomer)
	}
	return &controller

}
