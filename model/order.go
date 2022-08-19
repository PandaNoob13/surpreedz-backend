package model

import "time"

type Order struct {
	BaseModel       BaseModel     `gorm:"embedded" json:"baseModel"`
	ID              int           `json:"order_id" gorm:"primaryKey;not null"`
	// AccountId       int           `json:"buyer_id"`
	// Account         Account       `gorm:"foreignKey:AccountId" json:"account"`
	BuyerId       int           `json:"buyer_id" gorm:";not null"`
	Buyer         Account       `gorm:"foreignKey:BuyerId" json:"buyer"`
	ServiceDetailId int           `json:"service_detail_id" gorm:";not null"`
	ServiceDetail   ServiceDetail `gorm:"foreignKey:ServiceDetailId" json:"service_detail"`
	DueDate         time.Time     `json:"due_date" gorm:"-;not null"`
	TimeStamp       time.Time     `json:"time_stamp" gorm:"-;not null"`
	// OrderDate       time.Time     `json:"order_date" gorm:"-;not null"`
	// OrderStatus     []OrderStatus `gorm:"foreignKey:OrderId;references:ID"`
	// OrderRequest    OrderRequest  `gorm:"foreignKey:OrderId;references:ID"`
	// Feedback        Feedback      `gorm:"foreignKey:OrderId;references:ID"`
	// VideoResult     VideoResult   `gorm:"foreginKey:OrderId;references:ID"`
}

func (Order) TableName() string {
	return "mst_order"
}
