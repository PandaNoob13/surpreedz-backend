package model

import "gorm.io/gorm"

type VideoResult struct {
	gorm.Model
	ID        int    `json:"video_result_id" gorm:"primaryKey;not null"`
	OrderId   int    `json:"order_id" gorm:";not null"`
	VideoLink string `json:"video_link" gorm:"size:30;not null"`
	IsDeleted bool   `json:"is_deleted" gorm:";not null"`
}

func (VideoResult) TableName() string {
	return "mst_video_result"
}
