package utils

import (
	"context"
	"errors"
	"fmt"
	"surpreedz-backend/config"
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Token interface {
	CreateAccessToken(cred *dto.Credential) (*model.TokenDetails, error)
	VerifyAccessToken(tokenString string) (*model.AccessDetail, error)
	StoreAccessToken(userName string, tokenDetail *model.TokenDetails) error
	FetchAccessToken(accessDetail *model.AccessDetail) (string, error)
}

type token struct {
	cfg config.TokenConfig
}

// FetchAccessToken implements Token
func (t *token) FetchAccessToken(accessDetail *model.AccessDetail) (string, error) {
	if accessDetail != nil {
		result, err := t.cfg.Client.Get(context.Background(), accessDetail.AccessUuid).Result()
		if err != nil {
			return "", err
		}
		return result, nil
	} else {
		return "", errors.New("invalid access")
	}
}

func (t *token) StoreAccessToken(userName string, tokenDetail *model.TokenDetails) error {
	at := time.Unix(tokenDetail.AtExpires, 0)
	now := time.Now()
	err := t.cfg.Client.Set(context.Background(), tokenDetail.AccessUuid, userName, at.Sub(now)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (t *token) CreateAccessToken(cred *dto.Credential) (*model.TokenDetails, error) {
	td := &model.TokenDetails{}
	now := time.Now().UTC()
	end := now.Add(t.cfg.AccessTokenLifeTIme)

	td.AtExpires = end.Unix()
	td.AccessUuid = uuid.New().String()
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: t.cfg.ApplicationName,
		},
		Email: cred.Email,
		//Email:      cred.Email,
		AccessUUID: td.AccessUuid,
	}
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = end.Unix()
	token := jwt.NewWithClaims(
		t.cfg.JwtSingingMethod,
		claims,
	)
	newToken, err := token.SignedString([]byte(t.cfg.JwtSignatureKey))
	td.AccessToken = newToken
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (t *token) VerifyAccessToken(tokenString string) (*model.AccessDetail, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != t.cfg.JwtSingingMethod {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte(t.cfg.JwtSignatureKey), nil
	})
	if err != nil {
		fmt.Println("Error parsing token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != t.cfg.ApplicationName {
		fmt.Println("Token invalid...")
		return nil, err
	}

	fmt.Println("claims : ", claims)
	accessUuid := claims["AccessUUID"].(string)
	userName := claims["Email"].(string)
	return &model.AccessDetail{
		AccessUuid: accessUuid,
		Username:   userName,
	}, nil
}

func NewTokenService(cfg config.TokenConfig) Token {
	return &token{cfg: cfg}
}
