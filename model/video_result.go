package model

type VideoResult struct {
	BaseModel BaseModel `gorm:"embedded" json:"baseModel"`
	ID        int       `json:"video_result_id" gorm:"primaryKey;not null"`
	OrderId   int       `json:"order_id" gorm:";not null"`
	Order	Order 		`gorm:"foreignKey:OrderId" json:"order"`
	VideoLink string    `json:"video_link" gorm:"size:30;not null"`
	IsDeleted bool      `json:"is_deleted" gorm:"-;not null"`
}

func (VideoResult) TableName() string {
	return "mst_video_result"
}
