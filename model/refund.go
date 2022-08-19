package model

import "time"

type Refund struct {
	BaseModel     BaseModel   `gorm:"embedded" json:"baseModel"`
	ID            int         `json:"refund_id" gorm:"primaryKey;not null"`
	OrderStatusId int         `json:"order_status_id" gorm:";not null"`
	OrderStatus   OrderStatus `gorm:"foreignKey:OrderStatusId" json:"order_status"`
	Reason        string      `json:"reason" gorm:"size:50;not null"`
	TimeStamp     time.Time   `json:"time_stamp" gorm:"-;not null"`
}

func (Refund) TableName() string {
	return "mst_refund"
}
