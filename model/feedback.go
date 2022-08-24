package model

type Feedback struct {
	ID        int     `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	OrderId   int     `json:"order_id" gorm:";not null"`
	Review    string  `json:"review" gorm:";not null"`
	Rating    float32 `json:"rating" gorm:";not null"`
	IsDeleted bool    `json:"is_deleted" gorm:";not null"`
	Base_model
}

func (Feedback) TableName() string {
	return "mst_feedback"
}
