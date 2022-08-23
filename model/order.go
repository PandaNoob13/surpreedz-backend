package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID              int           `json:"order_id" gorm:"primaryKey;not null"`
	BuyerId         int           `json:"buyer_id" gorm:";not null"`
	ServiceDetailId int           `json:"service_detail_id" gorm:";not null"`
	DueDate         string        `json:"due_date" gorm:";not null"`
	OrderStatus     []OrderStatus `gorm:"foreignKey:OrderId;references:ID"`
	OrderRequest    OrderRequest  `gorm:"foreignKey:OrderId;references:ID"`
	Feedback        Feedback      `gorm:"foreignKey:OrderId;references:ID"`
	VideoResult     VideoResult   `gorm:"foreignKey:OrderId;references:ID"`
}

func (Order) TableName() string {
	return "mst_order"
}
