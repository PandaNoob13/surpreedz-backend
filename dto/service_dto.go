package dto

type ServiceDto struct {
	SellerId    int    `json:"seller_id"`
	Role        string `json:"role"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	VideoLink   string `json:"video_link"`
}
