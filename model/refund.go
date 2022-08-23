package model

import (
	"gorm.io/gorm"
)

type Refund struct {
	gorm.Model
	ID            int    `json:"refund_id" gorm:"primaryKey;not null"`
	OrderStatusId int    `json:"order_status_id" gorm:";not null"`
	Reason        string `json:"reason" gorm:"size:50;not null"`
}

func (Refund) TableName() string {
	return "mst_refund"
}
