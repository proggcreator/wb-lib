package repository

import (
	"strconv"
	"time"

	wbsql "git.wildberries.ru/finance/go-infrastructure/database/v2"
)

const (
	MaxOpen     int           = 10
	MaxIdle     int           = 5
	MaxLifetime time.Duration = time.Second * 10
)

type Config struct {
	Host              string
	Port              string
	Username          string
	Password          string
	DBName            string
	SSLMode           string
	ElacticHost       string
	ElacticAppName    string
	ElacticAppVersion string
}

func NewPostgresDB(cfg Config) (wbsql.DbConnecter, error) {
	MaxOpen := MaxOpen
	MaxIdle := MaxIdle
	MaxLifetime := MaxLifetime
	port, err := strconv.Atoi(cfg.Port)
	if err != nil {
		//err
	}

	db, err := wbsql.NewPgsqlSingleConnecter(wbsql.SqlConnectionConfig{
		Host:                  cfg.Host,
		Name:                  cfg.Username,
		User:                  cfg.Username,
		Password:              cfg.Password,
		Port:                  &port,
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
