package controller

import (
	"fmt"
	"net/http"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	router            *gin.Engine
	tokenService      utils.Token
	ucFindAccByEmail  usecase.FindAccountUseCase
	ucFindPassByAccId usecase.FindPasswordUseCase
}

func (l *LoginController) loginAkunCustomer(ctx *gin.Context) {

	var user dto.Credential
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	AccRes, dataUrl, err := l.ucFindAccByEmail.FindAccountByEmail(user.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}
	AccDtoRes := dto.AccountCreateDto{
		Account: AccRes,
	}
	AccDtoRes.StringJoinDate = AccRes.JoinDate.Format("2006-January-02")
	AccDtoRes.DataUrl = dataUrl
	PassRes, err := l.ucFindPassByAccId.FindPasswordByAccountId(AccRes.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(PassRes.Password), []byte(user.Password))

	if user.Email == AccRes.Email && err == nil {
		token, err := l.tokenService.CreateAccessToken(&user)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		err = l.tokenService.StoreAccessToken(user.Email, token)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"token":   token,
			"account": AccDtoRes,
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "email and password did not match",
		})
	}
}

func NewLoginController(router *gin.Engine, tokenService utils.Token, ucFindAccByEmail usecase.FindAccountUseCase, ucFindPassByAccId usecase.FindPasswordUseCase) *LoginController {

	controller := LoginController{
		router:            router,
		tokenService:      tokenService,
		ucFindAccByEmail:  ucFindAccByEmail,
		ucFindPassByAccId: ucFindPassByAccId,
	}

	rLogin := router.Group("/api")
	{
		rLogin.POST("/auth/login", controller.loginAkunCustomer)
	}
	return &controller

}
