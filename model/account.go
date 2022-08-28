package model

import (
	"time"
)

type Account struct {
	ID             int           `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Email          string        `json:"email" gorm:";not null"`
	AccountDetail  AccountDetail `gorm:"foreignKey:AccountId;references:ID"`
	JoinDate       time.Time     `json:"join_date" gorm:";not null"`
	StringJoinDate string        `json:"string_join_date"`
	Orders         []Order       `gorm:"foreignKey:BuyerId;references:ID"`
	ServiceDetail  ServiceDetail `gorm:"foreignKey:SellerId;references:ID"`
	DataUrl        string        `json:"data_url"`
	Base_model
}

func (Account) TableName() string {
	return "mst_account"
}
