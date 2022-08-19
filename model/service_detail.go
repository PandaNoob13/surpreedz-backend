package model

type ServiceDetail struct {
	BaseModel BaseModel `gorm:"embedded" json:"baseModel"`
	ID        int       `json:"service_detail_id" gorm:"primaryKey;not null"`
	SellerId  int       `json:"seller_id" gorm:";not null"`
	Seller    Account   `gorm:"foreignKey:SellerId" json:"seller"`
	// AccountId     int            `json:"seller_id"`
	// Account       Account        `gorm:"foreignKey:AccountId" json:"account"`
	Role      string    `json:"role" gorm:"size:20;not null"`
	Bio       string    `json:"bio" gorm:"size:30;not null"`
	// ServicePrice  ServicePrice   `gorm:"foreignKey:ServiceDetailId;references:ID"`
	// VideoProfiles []VideoProfile `gorm:"foreignKey:ServiceDetailId;references:ID"`
	// Orders        []Order        `gorm:"foreignKey:ServiceDetailId;references:ID"`
}

func (ServiceDetail) TableName() string {
	return "mst_service_detail"
}
