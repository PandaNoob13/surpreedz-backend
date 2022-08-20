package model

import "gorm.io/gorm"

type VideoProfile struct {
	gorm.Model
	ID               int    `json:"video_profile_id" gorm:"primaryKey;not null"`
	ServiceDetailId  int    `json:"service_detail_id" gorm:";not null"`
	VideoProfileLink string `json:"video_profile_link" gorm:"size:100;not null"`
	IsDeleted        bool   `json:"is_deleted" gorm:";not null"`
}

func (VideoProfile) TableName() string {
	return "mst_video_profile"
}
