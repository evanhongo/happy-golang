package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/evanhongo/happy-golang/api/httputil"
	"github.com/evanhongo/happy-golang/config"
	"github.com/evanhongo/happy-golang/pkg/lua"
	"github.com/evanhongo/happy-golang/pkg/util/custom_errors"
	service "github.com/evanhongo/happy-golang/service/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const (
	IPLimitPeriod     = 3600
	IPLimitTimeFormat = "2006-01-02 15:04:05"
	IPLimitMaximum    = 1000
)

type Middleware struct {
	service service.IAuthService
	redis   *redis.Client
}

func (m *Middleware) Auth(c *gin.Context) {
	cfg := config.GetConfig()
	userId, err := m.service.VerifyIdToken(c.GetHeader("Authorization"), cfg.JWT_SECRET)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &httputil.HttpErrorBody{Error: custom_errors.ErrNotAuthenticated.Error()})
		c.Abort()
		return
	}
	newToken, err := m.service.Sign(userId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			&httputil.HttpErrorBody{Error: custom_errors.ErrServerProblem.Error()},
		)
		c.Abort()
		return
	}
	// update JWT Token
	c.Header("Access-Control-Expose-Headers", "Authorization")
	c.Header("Authorization", newToken)
	c.Set("userId", userId)
	c.Next()
}

func (m *Middleware) RateLimit(c *gin.Context) {
	userId := c.GetString("userId")
	key := fmt.Sprintf("%s-%s-%s-%s", c.Request.URL.Path, c.Request.Method, c.ClientIP(), userId)
	now := time.Now().Unix()
	args := []interface{}{now, IPLimitMaximum, IPLimitPeriod}
	script := redis.NewScript(lua.RATE_LIMITER_SCRIPT)
	value, err := script.Run(m.redis, []string{key}, args...).Result()

	if err != nil {
		c.JSON(http.StatusInternalServerError, &httputil.HttpErrorBody{Error: err.Error()})
		c.Abort()
		return
	}

	result := value.([]interface{})
	remaining := result[0].(int64)
	if remaining == -1 {
		c.JSON(http.StatusTooManyRequests, nil)
		c.Abort()
		return
	}
	t := result[1].(int64)
	reset := time.Unix(t, 0).Format(IPLimitTimeFormat)

	c.Header("X-RateLimit-Remaining", strconv.FormatInt(remaining, 10))
	c.Header("X-RateLimit-Reset", reset)
	c.Next()
}

func NewMiddleware(service service.IAuthService, redis *redis.Client) *Middleware {
	return &Middleware{service: service, redis: redis}
}
