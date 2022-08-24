package dto

type OrderStatusDto struct {
	OrderId       int    `json:"order_id"`
	Status        string `json:"status"`
	ResonOfRefund string `json:"reason_of_refund"`
}
