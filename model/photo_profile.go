package model

type PhotoProfile struct {
	ID              int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	AccountDetailId int    `json:"account_detail_id" gorm:";not null"`
	PhotoLink       string `json:"photo_link" gorm:";not null"`
	IsDeleted       bool   `json:"is_deleted" gorm:";not null"`
	Base_model
}

func (PhotoProfile) TableName() string {
	return "mst_photo_profile"
}
