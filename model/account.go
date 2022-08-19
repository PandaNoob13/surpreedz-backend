package model

type Account struct {
	BaseModel BaseModel `gorm:"embedded" json:"baseModel"`
	ID        int       `json:"account_id" gorm:"primaryKey;not null"`
	Email     string    `json:"account_email" gorm:"size:100;not null"`
	Password  string    `json:"account_password" gorm:"size:20;not null"`
	// JoinDate      time.Time     `json:"join_date" gorm:"-;not null"`
	// AccountDetail AccountDetail `gorm:"foreignKey:AccountId;references:ID"`
	// Orders        []Order       `gorm:"foreignKey:BuyerId;references:ID"`
	// ServiceDetail ServiceDetail `gorm:"foreignKey:SellerId;references:ID"`
}

func (Account) TableName() string {
	return "mst_account"
}
