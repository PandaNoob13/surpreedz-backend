package dto

import "surpreedz-backend/model"

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

type AccountCreateDto struct {
	model.Account  `json:"account"`
	DataUrl        string `json:"data_url"`
	StringJoinDate string `json:"string_join_date"`

	model.Base_model
}

type PhotoVerifyForCMS struct {
	Email     string `json:"email"`
	PhotoLink string `json:"photo_link"`
}

type VerifyFromCMS struct {
	AccountId       int    `json:"account_id"`
	VerifiedStatus  bool   `json:"verified_status"`
	VerifiedRequest string `json:"verified_request"`
}

type EmailInput struct {
	Email string `json:"email"`
}

type AccIdInput struct {
	AccountId int `json:"account_id"`
}