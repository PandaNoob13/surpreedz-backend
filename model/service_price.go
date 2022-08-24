package model

import (
	"time"
)

type ServicePrice struct {
	ID              int       `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	ServiceDetailId int       `json:"service_detail_id" gorm:";not null"`
	Price           int       `json:"price" gorm:";not null"`
	Date            time.Time `json:"date" gorm:";not null"`
	Base_model
}

func (ServicePrice) TableName() string {
	return "mst_service_price"
}
