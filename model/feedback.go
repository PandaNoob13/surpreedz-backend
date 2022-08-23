package model

import "gorm.io/gorm"

type Feedback struct {
	ID        int     `json:"feedback_id" gorm:"primaryKey;not null"`
	OrderId   int     `json:"order_id" gorm:";not null"`
	Review    string  `json:"review" gorm:"size:100;not null"`
	Rating    float32 `json:"rating" gorm:";not null"`
	IsDeleted bool    `json:"is_deleted" gorm:";not null"`
	gorm.Model
}

func (Feedback) TableName() string {
	return "mst_feedback"
}
