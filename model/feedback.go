package model

type Feedback struct {
	BaseModel BaseModel `gorm:"embedded" json:"baseModel"`
	ID        int       `json:"feedback_id" gorm:"primaryKey;not null"`
	OrderId   int       `json:"order_id" gorm:";not null"`
	Order	Order 		`gorm:"foreignKey:OrderId" json:"order"`
	Review    string    `json:"review" gorm:"size:100;not null"`
	Rating    float32   `json:"rating" gorm:"-;not null"`
	IsDeleted bool      `json:"is_deleted" gorm:"-;not null"`
}

func (Feedback) TableName() string {
	return "mst_feedback"
}
