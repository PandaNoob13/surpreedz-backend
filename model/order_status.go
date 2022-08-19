package model

import "time"

type OrderStatus struct {
	BaseModel BaseModel `gorm:"embedded" json:"baseModel"`
	ID        int       `json:"order_status_id" gorm:"primaryKey;not null"`
	OrderId   int       `json:"order_id" gorm:";not null"`
	Order     Order     `gorm:"foreignKey:OrderId" json:"order"`
	Status    string    `json:"order_status" gorm:"size:15;not null"`
	TimeStamp time.Time `json:"time_stamp" gorm:"-;not null"`
	//Refund    Refund    `gorm:"foreignKey:OrderStatusId;references:ID"`
}

func (OrderStatus) TableName() string {
	return "mst_order_status"
}
