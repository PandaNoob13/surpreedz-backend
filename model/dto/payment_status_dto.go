package dto

type PaymentStatusDto struct {
	OrderId       string `json:"order_id"`
	StatusPayment string `json:"transaction_status"`
	PaymentType   string `json:"payment_type"`
}
