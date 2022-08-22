package utils

import "github.com/golang-jwt/jwt"

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"Email"`
	//Email      string `json:"Email"`
	AccessUUID string `json:"AccessUUID"`
}
