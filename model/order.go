package model

type Order struct {
	ID              int           `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	BuyerId         int           `json:"buyer_id" gorm:";not null"`
	ServiceDetailId int           `json:"service_detail_id" gorm:";not null"`
	DueDate         string        `json:"due_date" gorm:";not null"`
	OrderStatus     []OrderStatus `gorm:"foreignKey:OrderId;references:ID"`
	OrderRequest    OrderRequest  `gorm:"foreignKey:OrderId;references:ID"`
	Feedback        Feedback      `gorm:"foreignKey:OrderId;references:ID"`
	VideoResult     VideoResult   `gorm:"foreignKey:OrderId;references:ID"`
	Base_model
}

func (Order) TableName() string {
	return "mst_order"
}
