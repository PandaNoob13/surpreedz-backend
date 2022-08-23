package model

import "gorm.io/gorm"

type OrderRequest struct {
	ID            int    `json:"order_request_id" gorm:"primaryKey;not null"`
	OrderId       int    `json:"order_id" gorm:";not null"`
	Ocassion      string `json:"ocassoion" gorm:"size:30;not null"`
	RecipientName string `json:"recipient_name" gorm:"size:20;not null"`
	Message       string `json:"message" gorm:"size:100;not null"`
	Description   string `json:"description" gorm:"size:80;not null"`
	gorm.Model
}

func (OrderRequest) TableName() string {
	return "mst_order_request"
}
