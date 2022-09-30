package model

import "time"

type PaymentStatus struct {
	ID            int       `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	OrderId       string    `json:"order_id" gorm:";not null"`
	StatusPayment string    `json:"transaction_status" gorm:";not null"`
	PaymentType   string    `json:"payment_type" gorm:";not null"`
	TimeUpdated   time.Time `json:"time_updated" gorm:";not null"`
	Base_model
}

func (PaymentStatus) TableName() string {
	return "mst_payment_status"
}
