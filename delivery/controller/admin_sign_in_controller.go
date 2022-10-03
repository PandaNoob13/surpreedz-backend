package controller

import (
	"fmt"
	"net/http"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
)

type AdminLoginController struct {
	router         *gin.Engine
	tokenService   utils.Token
	ucFindAdminAcc usecase.FindAdminAccUseCase
}

func (a *AdminLoginController) loginAkunAdmin(ctx *gin.Context) {

	var user dto.Credential
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	AccRes, err := a.ucFindAdminAcc.FindAdminAccByUsername(user.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}

	if user.Email == AccRes.Username && user.Password == AccRes.Password {
		token, err := a.tokenService.CreateAccessToken(&user)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		err = a.tokenService.StoreAccessToken(user.Email, token)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"token":   token,
			"account": AccRes.Username,
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "email and password did not match",
		})
	}
}

func NewAdminLoginController(router *gin.Engine, tokenService utils.Token, ucFindAdminAcc usecase.FindAdminAccUseCase) *AdminLoginController {

	controller := AdminLoginController{
		router:         router,
		tokenService:   tokenService,
		ucFindAdminAcc: ucFindAdminAcc,
	}

	rLoginAdmin := router.Group("/cms")
	{
		rLoginAdmin.POST("/auth/admin/login", controller.loginAkunAdmin)
	}
	return &controller

}
