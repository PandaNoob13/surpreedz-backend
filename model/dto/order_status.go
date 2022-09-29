package dto

type OrderStatusDto struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}
