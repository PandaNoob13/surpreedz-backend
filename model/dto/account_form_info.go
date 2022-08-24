package dto

type AccountFormInfo struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	PhotoLink string `json:"photo_link"`
	IsDeleted bool   `json:"is_deleted"`
}
