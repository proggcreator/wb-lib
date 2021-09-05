package repository

import (
	wbsql "git.wildberries.ru/finance/go-infrastructure/database/v2"

	logs "git.wildberries.ru/finance/go-infrastructure/elasticlog"
	restful "github.com/proggcreator/wb-lib"
)

type EmplWork interface {
	CreateEmpl(employee restful.Employee) (string, error)
	GetAllEmpl() ([]restful.Employee, error)
	GetByIdEmpl(userId string) (restful.Employee, error)
	DeleteEmpl(userId string) error
	UpdateEmpl(newemployee restful.Employee) error
}

type Repository struct {
	EmplWork
}

func NewRepository(db wbsql.DbConnecter, wblogger *logs.Logger) *Repository {
	return &Repository{
		EmplWork: NewEmplWorkPostgres(db, wblogger),
	}
}
