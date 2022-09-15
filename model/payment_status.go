package model

import "time"

type PaymentStatus struct {
	ID            int       `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	OrderId       int       `json:"order_id" gorm:";not null"`
	StatusPayment string    `json:"payment_status" gorm:";not null"`
	TimeUpdated   time.Time `json:"time_updated" gorm:";not null"`
	Base_model
}

func (PaymentStatus) TableName() string {
	return "mst_payment_status"
}
