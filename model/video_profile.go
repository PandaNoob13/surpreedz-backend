package model

type VideoProfile struct {
	BaseModel        BaseModel     `gorm:"embedded" json:"baseModel"`
	ID               int           `json:"video_profile_id" gorm:"primaryKey;not null"`
	ServiceDetailId  int           `json:"service_detail_id" gorm:";not null"`
	ServiceDetail    ServiceDetail `gorm:"foreignKey:ServiceDetailId" json:"service_detail"`
	VideoProfileLink string        `json:"video_profile_link" gorm:"size:30;not null"`
	IsDeleted        bool          `json:"is_deleted" gorm:"false;not null"`
}

func (VideoProfile) TableName() string {
	return "mst_video_profile"
}
