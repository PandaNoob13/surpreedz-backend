package model

type AccountDetail struct {
	ID             int            `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	AccountId      int            `json:"account_id" gorm:";not null"`
	Name           string         `json:"name" gorm:";not null"`
	Location       string         `json:"location" gorm:";not null"`
	PhotoProfiles  []PhotoProfile `gorm:"foreignKey:AccountDetailId;references:ID"`
	VerifiedStatus bool           `json:"verified_status" gorm:"default:false"`
	Base_model
}

func (AccountDetail) TableName() string {
	return "mst_account_detail"
}
