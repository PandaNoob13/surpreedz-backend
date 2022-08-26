package dto

type PaymentInput struct {
	OrderId string `json:"order_id"`
	Amount  int    `json:"amount"`
}
