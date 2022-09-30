package model

type OrderRequest struct {
	ID            int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	OrderId       string    `json:"order_id" gorm:";not null"`
	Occasion      string `json:"occasion" gorm:";not null"`
	RecipientName string `json:"recipient_name" gorm:";not null"`
	Message       string `json:"message" gorm:";not null"`
	Description   string `json:"description" gorm:";not null"`
	Base_model
}

func (OrderRequest) TableName() string {
	return "mst_order_request"
}
