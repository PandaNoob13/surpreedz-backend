package controller

import (
	"net/http"
	"surpreedz-backend/model/dto"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentController struct {
	router *gin.Engine
}

func (p *PaymentController) doPayment(ctx *gin.Context) {

	var input dto.PaymentInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	// 1. Initiate Snap client
	var s snap.Client
	s.New("Mid-server-c8LkxQRU7RvCQDyq4kWEBCj3", midtrans.Sandbox)
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
			FName: "Samuel",
			LName: "Maynard",
			Email: "sam@may.com",
			Phone: "081234567890",
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	ctx.JSON(http.StatusOK, gin.H{
		"token":        snapResp.Token,
		"redirect_url": snapResp.RedirectURL,
	})

}

func NewPaymentController(router *gin.Engine) *PaymentController {

	controller := PaymentController{
		router: router,
	}

	rPayment := router.Group("/order")
	{
		rPayment.POST("/payment", controller.doPayment)
	}
	return &controller

}
