package model

import "gorm.io/gorm"

type AccountDetail struct {
	ID            int            `json:"account_detail_id" gorm:"primaryKey;not null"`
	AccountId     int            `json:"account_id" gorm:";not null"`
	UserName      string         `json:"username" gorm:"size:20;not null"`
	Location      string         `json:"location" gorm:"size:15;not null"`
	PhotoProfiles []PhotoProfile `gorm:"foreignKey:AccountDetailId;references:ID"`
	gorm.Model
}

func (AccountDetail) TableName() string {
	return "mst_account_detail"
}
