package dto

type PaymentStatusDto struct {
	OrderId       int    `json:"order_id"`
	StatusPayment string `json:"transaction_status"`
	PaymentType   string `json:"payment_type"`
}
