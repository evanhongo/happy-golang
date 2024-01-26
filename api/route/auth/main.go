package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/evanhongo/happy-golang/internal/env"
	service "github.com/evanhongo/happy-golang/service/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.IAuthService
}

func (handler *AuthHandler) RetrieveAuthorizationCode(c *gin.Context) {
	env := env.GetEnv()
	url := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=https://www.googleapis.com/auth/userinfo.profile openid profile email", env.GOOGLE_CLIENT_ID, env.GOOGLE_REDIRECT_URL)
	c.Redirect(http.StatusSeeOther, url)
}

func (handler *AuthHandler) RetrieveAccessToken(c *gin.Context) {
	env := env.GetEnv()
	authorizationCode := c.Query("code")
	tokenEndpoint := fmt.Sprintf("https://oauth2.googleapis.com/token?code=%s&client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s", authorizationCode, env.GOOGLE_CLIENT_ID, env.GOOGLE_CLIENT_SECRET, env.GOOGLE_REDIRECT_URL)
	tokenResponse := struct {
		IdToken     string `json:"id_token"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}{}

	if resp, err := http.Post(tokenEndpoint, "text/plain", bytes.NewReader([]byte(``))); err != nil { // get token
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/ping")
	} else if body, err := io.ReadAll(resp.Body); err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/ping")
	} else if err := json.Unmarshal(body, &tokenResponse); err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/ping")
	} else if err := handler.service.VerifyIdToken(tokenResponse.IdToken); err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/ping")
	}

	c.Redirect(http.StatusFound, "/")
}

func NewAuthHandler(service service.IAuthService) *AuthHandler {
	return &AuthHandler{
		service,
	}
}
