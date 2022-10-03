package dto

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
