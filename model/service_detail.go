package model

type ServiceDetail struct {
	ID            int            `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	SellerId      int            `json:"seller_id" gorm:";not null"`
	Role          string         `json:"role" gorm:";not null"`
	Description   string         `json:"description" gorm:";not null"`
	ServicePrices []ServicePrice `gorm:"foreignKey:ServiceDetailId;references:ID"`
	VideoProfiles []VideoProfile `gorm:"foreignKey:ServiceDetailId;references:ID"`
	Orders        []Order        `gorm:"foreignKey:ServiceDetailId;references:ID"`
	Base_model
}

func (ServiceDetail) TableName() string {
	return "mst_service_detail"
}
