package model

type ServiceDetail struct {
	ID            int            `json:"service_detail_id" gorm:"primaryKey;not null"`
	SellerId      int            `json:"seller_id" gorm:";not null"`
	Role          string         `json:"role" gorm:"size:20;not null"`
	Description   string         `json:"description" gorm:"size:30;not null"`
	ServicePrices []ServicePrice `gorm:"foreignKey:ServiceDetailId;references:ID"`
	VideoProfiles []VideoProfile `gorm:"foreignKey:ServiceDetailId;references:ID"`
	Orders        []Order        `gorm:"foreignKey:ServiceDetailId;references:ID"`
}

func (ServiceDetail) TableName() string {
	return "mst_service_detail"
}
