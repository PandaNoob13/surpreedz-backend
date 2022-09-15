package dto

type OrderDto struct {
	BuyerId              int    `json:"buyer_id"`
	ServiceDetailId      int    `json:"service_detail_id"`
	DueDate              string `json:"due_date"`
	Occasion             string `json:"occasion"`
	RecipientName        string `json:"recipient_name"`
	Message              string `json:"message_to_recipient"`
	RecipientDescription string `json:"recipient_description"`
	StatusPayment        string `json:"payment_status"`
}
