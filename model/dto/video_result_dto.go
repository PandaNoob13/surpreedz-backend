package dto

import "surpreedz-backend/model"

type VideoResultDto struct {
	OrderId   int    `json:"order_id"`
	VideoLink string `json:"video_link"`
	DataUrl   string `json:"data_url"`
	model.Base_model
}
