package model

import "gorm.io/gorm"

type PhotoProfile struct {
	gorm.Model
	ID              int    `json:"photo_profile_id" gorm:"primaryKey;not null"`
	AccountDetailId int    `json:"account_detail_id" gorm:";not null"`
	PhotoLink       string `json:"photo_link" gorm:"size:30;not null"`
	IsDeleted       bool   `json:"is_deleted" gorm:";not null"`
}

func (PhotoProfile) TableName() string {
	return "mst_photo_profile"
}
