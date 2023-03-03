package env

import (
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once
var env *Env

type Env struct {
	ENVIRONMENT string `env:"ENVIRONMENT"`
	LOG_LEVEL   string `env:"LOG_LEVEL"`
	PORT        string `env:"PORT"`

	GOOGLE_CLIENT_ID     string `env:"GOOGLE_CLIENT_ID"`
	GOOGLE_CLIENT_SECRET string `env:"GOOGLE_CLIENT_SECRET"`
	GOOGLE_REDIRECT_URL  string `env:"GOOGLE_REDIRECT_URL"`

	JOB_QUEUE_BROKER           string `env:"JOB_QUEUE_BROKER"`
	JOB_QUEUE_RESULT_BACKEND   string `env:"JOB_QUEUE_RESULT_BACKEND"`
	JOB_QUEUE_RESULT_EXPIRE_IN int    `env:"JOB_QUEUE_RESULT_EXPIRE_IN "`
	JOB_QUEUE_DEFAULT_QUEUE    string `env:"JOB_QUEUE_DEFAULT_QUEUE"`
	JOB_QUEUE_WORKER_NUM       int    `env:"JOB_QUEUE_WORKER_NUM"`
	JOB_QUEUE_LOG_LEVEL        string `env:"JOB_QUEUE_LOG_LEVEL"`
}

func GetEnv() *Env {
	once.Do(func() {
		v := viper.New()
		v.SetConfigFile(".env")
		v.ReadInConfig()
		v.AutomaticEnv()
		v.Unmarshal(&env)
	})

	return env
}
