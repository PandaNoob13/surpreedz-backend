package dto

type EditServiceDto struct {
	Role        string `json:"role"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	VideoLink   string `json:"video_link"`
}
