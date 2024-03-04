package service

import "github.com/evanhongo/happy-golang/entity"

type GooglePublicKey struct {
	Kid string `json:"kid"`
	N   string `json:"n"`
}

type IAuthService interface {
	RegisterByEmail(input *struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}) error
	LoginByEmail(input *struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}) (*entity.User, error)
	GetGooglePublicKey(idToken string) (string, error)
	Sign(userId string) (token string, err error)
	VerifyIdToken(idToken string, key string) (userId string, err error)
}
