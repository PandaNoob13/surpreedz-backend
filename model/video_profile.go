package model

type VideoProfile struct {
	ID               int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	ServiceDetailId  int    `json:"service_detail_id" gorm:";not null"`
	VideoProfileLink string `json:"video_profile_link" gorm:";not null"`
	IsDeleted        bool   `json:"is_deleted" gorm:";not null"`
	Base_model
}

func (VideoProfile) TableName() string {
	return "mst_video_profile"
}
