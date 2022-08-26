package model

type Password struct {
	ID        int     `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	AccountId int     `json:"account_id" gorm:";not null"`
	Account   Account `gorm:"foreignKey:AccountId"`
	Password  string  `json:"password" gorm:";not null"`
	Base_model
}

func (Password) TableName() string {
	return "mst_password"
}
