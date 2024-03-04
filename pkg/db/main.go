package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/evanhongo/happy-golang/config"
)

func NewDb() (*gorm.DB, error) {
	cfg := config.GetConfig()
	connectStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME,
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
	)
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  connectStr,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}),
		&gorm.Config{TranslateError: true},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
