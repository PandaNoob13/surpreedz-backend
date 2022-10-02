package model

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Username string `json:"username" gorm:";not null"`
	Password string `json:"password" gorm:";not null"`
}

func (Admin) TableName() string {
	return "mst_admin"
}
