package service

import (
	"encoding/json"
	"errors"
	"time"

	"io"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/evanhongo/happy-golang/config"
	"github.com/evanhongo/happy-golang/entity"
	"github.com/evanhongo/happy-golang/pkg/logger"
	"github.com/evanhongo/happy-golang/pkg/util/custom_errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func (service *AuthService) RegisterByEmail(input *struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}) error {
	var user entity.User
	user.Name = input.Name
	user.Email = input.Email

	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	user.ID = uid.String()

	digest, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(digest)

	result := service.db.Exec(
		`insert into users(id, email, password_digest, name)
		select ?, ?, ?, ?
		where not exists (select 1 from users where email = ?)`,
		user.ID, user.Email, user.PasswordDigest, user.Name, user.Email,
	)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return custom_errors.ErrDuplicatedEmail
	}

	return nil
}

func (service *AuthService) LoginByEmail(input *struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}) (*entity.User, error) {
	var user = &entity.User{}
	err := service.db.Where("email = ?", input.Email).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, custom_errors.ErrEmailNotFound
		}
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(input.Password)) != nil {
		return nil, custom_errors.ErrWrongPassword
	}

	return user, nil
}

func (service *AuthService) GetGooglePublicKey(idToken string) (string, error) {
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
		return "", err
	} else if body, err := io.ReadAll(resp.Body); err != nil {
		return "", err
	} else if err = json.Unmarshal(body, &googleOpenIdConfig); err != nil {
		return "", err
	} else if resp, err = http.Get(googleOpenIdConfig.JwksURI); err != nil { // 取得公鑰
		return "", err
	} else if body, err = io.ReadAll(resp.Body); err != nil {
		return "", err
	} else if err = json.Unmarshal(body, &googleCert); err != nil {
		return "", err
	} else {
		data, err := jwt.Parse(idToken, nil)
		if err != nil {
			return "", err
		}

		kid := data.Header["kid"].(string) // 取得公鑰Id
		var publicKey string
		for _, k := range googleCert.Keys { // 從鑰匙串中找出 Id 正確的鑰匙
			if k.Kid == kid {
				publicKey = k.N
				break
			}
		}

		return publicKey, nil
	}
}

func (s *AuthService) Sign(userId string) (string, error) {
	cfg := config.GetConfig()
	claims := jwt.StandardClaims{
		Id:        userId,
		ExpiresAt: time.Now().Add(time.Duration(60 * time.Minute)).Unix(),
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(cfg.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) VerifyIdToken(idToken string, key string) (string, error) {
	tokenClaims, err := jwt.ParseWithClaims(idToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				err = errors.New("token is malformed")
			} else if v.Errors&jwt.ValidationErrorUnverifiable != 0 {
				err = errors.New("token could not be verified because of signing problems")
			} else if v.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				err = errors.New("signature validation failed")
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				err = errors.New("token is expired")
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				err = errors.New("token is not yet valid before sometime")
			} else {
				err = errors.New("can not handle this token")
			}
		}
		logger.Error(err.Error())
		return "", err
	}

	claims, _ := tokenClaims.Claims.(*jwt.StandardClaims)
	if claims.Id == "" {
		return "", errors.New("token is improper")
	}
	return claims.Id, err
}

func NewAuthService(db *gorm.DB) IAuthService {
	return &AuthService{db: db}
}
