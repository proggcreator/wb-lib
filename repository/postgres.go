package repository

import (
	"time"

	wbsql "git.wildberries.ru/finance/go-infrastructure/database/v2"
)

const (
	MaxOpen     int           = 10
	MaxIdle     int           = 5
	MaxLifetime time.Duration = time.Second * 10
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (wbsql.DbConnecter, error) {
	MaxOpen := MaxOpen
	MaxIdle := MaxIdle
	MaxLifetime := MaxLifetime

	db, err := wbsql.NewPgsqlSingleConnecter(wbsql.SqlConnectionConfig{
		Host:                  cfg.Host,
		Name:                  cfg.Username,
		User:                  cfg.Username,
		Password:              cfg.Password,
		Port:                  &cfg.Port,
		MaxOpenConnections:    &MaxOpen,
		MaxIdleConnections:    &MaxIdle,
		MaxConnectionLifetime: &MaxLifetime,
	})
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
