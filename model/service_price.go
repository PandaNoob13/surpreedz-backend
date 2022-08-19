package model

import (
	"time"
)

type ServicePrice struct {
	BaseModel       BaseModel     `gorm:"embedded" json:"baseModel"`
	ID              int           `json:"service_price_id" gorm:"primaryKey;not null"`
	ServiceDetailId int           `json:"service_detail_id" gorm:";not null"`
	ServiceDetail   ServiceDetail `gorm:"foreignKey:ServiceDetailId" json:"service_detail"`
	Price           int           `json:"price" gorm:"-;not null"`
	TimeStamp       time.Time     `json:"time_stamp" gorm:"-;not null"`
}

func (ServicePrice) TableName() string {
	return "mst_service_price"
}
