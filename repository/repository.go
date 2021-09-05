package repository

import (
	wbsql "git.wildberries.ru/finance/go-infrastructure/database/v2"

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

func NewRepository(db wbsql.DbConnecter) *Repository {
	return &Repository{
		EmplWork: NewEmplWorkPostgres(db),
	}
}
