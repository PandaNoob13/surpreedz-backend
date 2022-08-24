package model

import (
	"time"
)

type OrderStatus struct {
	ID      int       `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	OrderId int       `json:"order_id" gorm:";not null"`
	Status  string    `json:"order_status" gorm:";not null"`
	Date    time.Time `json:"date" gorm:";not null"`
	Refund  Refund    `gorm:"foreignKey:OrderStatusId;references:ID"`
	Base_model
}

func (OrderStatus) TableName() string {
	return "mst_order_status"
}
