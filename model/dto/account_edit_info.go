package dto

type AccountEditInfo struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	PhotoLink   string `json:"photo_link"`
	IsDeleted   bool   `json:"is_deleted"`
}

type EditProfileDto struct {
	AccountId int    `json:"account_id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	PhotoName string `json:"photo_name"`
	PhotoUrl  string `json:"url"`
	DataUrl   string `json:"data_url"`
}

type EditPasswordDto struct {
	AccountId   int    `json:"account_id"`
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
}
