package config

import (
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once
var cfg *Config

type Config struct {
	ENVIRONMENT string `env:"ENVIRONMENT"`
	LOG_LEVEL   string `env:"LOG_LEVEL"`
	PORT        string `env:"PORT"`

	DB_HOST     string `env:"DB_HOST"`
	DB_PORT     string `env:"DB_PORT"`
	DB_NAME     string `env:"DB_NAME"`
	DB_USERNAME string `env:"DB_USERNAME"`
	DB_PASSWORD string `env:"DB_PASSWORD"`

	REDIS_ENDPOINT string `env:"REDIS_ENDPOINT"`

	JWT_SECRET string `env:"JWT_SECRET"`

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

func GetConfig() *Config {
	once.Do(func() {
		v := viper.New()
		v.SetConfigFile(".env")
		v.ReadInConfig()
		v.AutomaticEnv()
		v.Unmarshal(&cfg)
	})

	return cfg
}
