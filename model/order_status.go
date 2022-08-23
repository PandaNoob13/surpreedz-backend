package model

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus struct {
	ID      int       `json:"order_status_id" gorm:"primaryKey;not null"`
	OrderId int       `json:"order_id" gorm:";not null"`
	Status  string    `json:"order_status" gorm:"size:15;not null"`
	Date    time.Time `json:"date" gorm:";not null"`
	Refund  Refund    `gorm:"foreignKey:OrderStatusId;references:ID"`
	gorm.Model
}

func (OrderStatus) TableName() string {
	return "mst_order_status"
}