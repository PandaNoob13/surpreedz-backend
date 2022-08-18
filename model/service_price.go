package model

import (
	"time"

	"gorm.io/gorm"
)

type ServicePrice struct {
	gorm.Model
	ID              int       `json:"service_price_id" gorm:"primaryKey;not null"`
	ServiceDetailId int       `json:"service_detail_id" gorm:"-;not null"`
	Price           int       `json:"price" gorm:"-;not null"`
	Date            time.Time `json:"date" gorm:"-;not null"`
}

func (ServicePrice) TableName() string {
	return "mst_service_price"
}
