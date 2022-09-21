package dto

type PaymentInput struct {
	Email   string `json:"email"`
	OrderId string `json:"order_id"`
	Amount  int    `json:"amount"`
}
