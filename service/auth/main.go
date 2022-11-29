package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/evanhongo/happy-golang/pkg/logger"
)

type AuthService struct {
}

func (service *AuthService) GetGooglePublicKeyChain() ([]GooglePublicKey, error) {
	const (
		googleOpenIdConfigURI = "https://accounts.google.com/.well-known/openid-configuration" // google Config URI
	)
	googleOpenIdConfig := struct {
		Issuer  string `json:"issuer"`
		JwksURI string `json:"jwks_uri"`
	}{}
	googleCert := struct {
		Keys []GooglePublicKey `json:"keys"`
	}{}

	if resp, err := http.Get(googleOpenIdConfigURI); err != nil { // 取得 Google OpenId Config
		return nil, err
	} else if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else if err = json.Unmarshal(body, &googleOpenIdConfig); err != nil {
		return nil, err
	} else if resp, err = http.Get(googleOpenIdConfig.JwksURI); err != nil { // 取得公鑰
		return nil, err
	} else if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else if err = json.Unmarshal(body, &googleCert); err != nil {
		return nil, err
	} else {
		return googleCert.Keys, nil
	}
}

func (s *AuthService) VerifyIdToken(idToken string) error {
	data, _ := jwt.Parse(idToken, nil)
	kid := data.Header["kid"].(string) // 取得公鑰Id

	var publicKey string
	googlePublicKeys, err := s.GetGooglePublicKeyChain()
	if err != nil {
		logger.Error("Unable to get google public key")
		return err
	}

	for _, k := range googlePublicKeys { // 從鑰匙串中找出 Id 正確的鑰匙
		if k.Kid == kid {
			publicKey = k.N
			break
		}
	}

	token, err := jwt.ParseWithClaims(idToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})

	if v, ok := err.(*jwt.ValidationError); ok && v.Errors == jwt.ValidationErrorMalformed {
		logger.Error("Token is malformed")
		err = errors.New("token is malformed")
	} else if claim, ok := token.Claims.(*jwt.StandardClaims); !ok {
		logger.Error("Unable to get claims")
		err = errors.New("unable to get claims")
	} else if time.Now().Unix() > claim.ExpiresAt {
		logger.Error("Token expired")
		err = errors.New("uoken expired")
	} else if claim.Issuer != "https://accounts.google.com" && claim.Issuer != "accounts.google.com" {
		logger.Error("Invalid Issuer")
		err = errors.New("invalid Issuer")
	} else if claim.Audience != "599903054173-b46j6a9lq4pbgsbe10ctkrtnr0drob9s.apps.googleusercontent.com" {
		logger.Error("Invalid Audience")
		err = errors.New("invalid audience")
	} else {
		logger.Info("ID Token Validation Success")
		err = nil
	}
	return err
}

func NewAuthService() IAuthService {
	return &AuthService{}
}
