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
