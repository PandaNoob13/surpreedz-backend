package model

import (
	"time"
)

type Account struct {
	ID            int           `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Email         string        `json:"email" gorm:";not null"`
	//Password      string        `json:"password" gorm:";not null"`
	AccountDetail AccountDetail `gorm:"foreignKey:AccountId;references:ID"`
	JoinDate      time.Time     `json:"join_date" gorm:";not null"`

	Orders        []Order       `gorm:"foreignKey:BuyerId;references:ID"`
	ServiceDetail ServiceDetail `gorm:"foreignKey:SellerId;references:ID"`
	Base_model
}

func (Account) TableName() string {
	return "mst_account"
}
