package auth_route

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/evanhongo/happy-golang/api/httputil"
	"github.com/evanhongo/happy-golang/config"
	"github.com/evanhongo/happy-golang/pkg/logger"
	"github.com/evanhongo/happy-golang/pkg/util/custom_errors"
	service "github.com/evanhongo/happy-golang/service/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.IAuthService
}

func (handler *AuthHandler) RegisterByEmail(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, &httputil.HttpErrorBody{Error: err.Error()})
		return
	}

	if err := handler.service.RegisterByEmail(&req); err != nil {
		logger.Error(err)
		if errors.Is(err, custom_errors.ErrDuplicatedEmail) {
			c.JSON(http.StatusBadRequest, &httputil.HttpErrorBody{Error: err.Error()})
			return
		}
		c.JSON(
			http.StatusInternalServerError,
			&httputil.HttpErrorBody{Error: custom_errors.ErrServerProblem.Error()},
		)
		return
	}

	c.Status(http.StatusCreated)
}

func (handler *AuthHandler) LoginByEmail(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, &httputil.HttpErrorBody{Error: err.Error()})
		return
	}

	user, err := handler.service.LoginByEmail(&req)
	if err != nil {
		logger.Error(err)
		if errors.Is(err, custom_errors.ErrEmailNotFound) ||
			errors.Is(err, custom_errors.ErrWrongPassword) {
			c.JSON(http.StatusBadRequest, &httputil.HttpErrorBody{Error: err.Error()})
			return
		}
		c.JSON(
			http.StatusInternalServerError,
			&httputil.HttpErrorBody{Error: custom_errors.ErrServerProblem.Error()},
		)
		return
	}
	token, err := handler.service.Sign(user.ID)
	if err != nil {
		logger.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			&httputil.HttpErrorBody{Error: custom_errors.ErrServerProblem.Error()},
		)
		return
	}
	c.Header("Access-Control-Expose-Headers", "Authorization")
	c.Header("Authorization", token)
	c.Status(http.StatusFound)
}

func (handler *AuthHandler) RetrieveAuthorizationCode(c *gin.Context) {
	cfg := config.GetConfig()
	url := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=https://www.googleapis.com/auth/userinfo.profile openid profile email", cfg.GOOGLE_CLIENT_ID, cfg.GOOGLE_REDIRECT_URL)
	c.Redirect(http.StatusSeeOther, url)
}

func (handler *AuthHandler) RetrieveAccessToken(c *gin.Context) {
	cfg := config.GetConfig()
	authorizationCode := c.Query("code")
	tokenEndpoint := fmt.Sprintf("https://oauth2.googleapis.com/token?code=%s&client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s", authorizationCode, cfg.GOOGLE_CLIENT_ID, cfg.GOOGLE_CLIENT_SECRET, cfg.GOOGLE_REDIRECT_URL)
	tokenResponse := struct {
		IdToken     string `json:"id_token"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}{}

	if resp, err := http.Post(tokenEndpoint, "text/plain", bytes.NewReader([]byte(``))); err != nil { // get token
		fmt.Println(err)
		c.Redirect(http.StatusInternalServerError, "/health/ping")
	} else if body, err := io.ReadAll(resp.Body); err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusInternalServerError, "/health/ping")
	} else if err := json.Unmarshal(body, &tokenResponse); err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusInternalServerError, "/health/ping")
	} else if key, err := handler.service.GetGooglePublicKey(tokenResponse.IdToken); err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusInternalServerError, "/health/ping")
	} else if _, err := handler.service.VerifyIdToken(tokenResponse.IdToken, key); err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusInternalServerError, "/health/ping")
	} else {
		c.Header("Access-Control-Expose-Headers", "Authorization")
		c.Header("Authorization", tokenResponse.IdToken)
		c.Redirect(http.StatusFound, "/health/ping")
	}
}

func NewAuthHandler(service service.IAuthService) *AuthHandler {
	return &AuthHandler{
		service,
	}
}
