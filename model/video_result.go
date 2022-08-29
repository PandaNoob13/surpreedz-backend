package model

type VideoResult struct {
	ID        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	OrderId   int    `json:"order_id" gorm:";not null"`
	VideoLink string `json:"video_link" gorm:";not null"`
	IsDeleted bool   `json:"is_deleted" gorm:";not null"`
	DataUrl   string `json:"data_url"`
	Base_model
}

func (VideoResult) TableName() string {
	return "mst_video_result"
}
