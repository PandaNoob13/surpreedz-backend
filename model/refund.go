package model

type Refund struct {
	ID            int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	OrderStatusId int    `json:"order_status_id" gorm:";not null"`
	Reason        string `json:"reason" gorm:";not null"`
	//Date          time.Time `json:"date" gorm:";not null"`
	Base_model
}

func (Refund) TableName() string {
	return "mst_refund"
}
