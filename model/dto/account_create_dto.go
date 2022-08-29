package dto

import (
	"surpreedz-backend/model"
)

type AccountCreateDto struct {
	model.Account  `json:"account"`
	DataUrl        string `json:"data_url"`
	StringJoinDate string `json:"string_join_date"`

	model.Base_model
}
