package controller

import (
	"net/http"
	"os"
	"surpreedz-backend/model/dto"
	"surpreedz-backend/usecase"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentController struct {
	router           *gin.Engine
	ucFindAccByEmail usecase.FindAccountUseCase
}

func (p *PaymentController) doPayment(ctx *gin.Context) {

	var input dto.PaymentInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	AccRes, _, err := p.ucFindAccByEmail.FindAccountByEmail(input.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}

	// 1. Initiate Snap client
	var s snap.Client
	midtransServerKey := os.Getenv("MIDTRANS_SERVER_KEY")
	s.New(midtransServerKey, midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  input.OrderId,
			GrossAmt: int64(input.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: AccRes.AccountDetail.Name,
			Email: AccRes.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	ctx.JSON(http.StatusOK, gin.H{
		"token":        snapResp.Token,
		"redirect_url": snapResp.RedirectURL,
	})

}

func NewPaymentController(router *gin.Engine, ucFindAccByEmail usecase.FindAccountUseCase) *PaymentController {

	controller := PaymentController{
		router:           router,
		ucFindAccByEmail: ucFindAccByEmail,
	}

	rPayment := router.Group("/order")
	{
		rPayment.POST("/payment", controller.doPayment)
	}
	return &controller

}
