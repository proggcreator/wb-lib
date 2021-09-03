package repository

import (
	"context"
	"database/sql"

	restful "github.com/proggcreator/wb-lib"
)

type EmplWork interface {
	CreateEmpl(empl restful.Employee, ctx context.Context) (string, error)
	GetAllEmpl(ctx context.Context) ([]restful.Employee, error)
	GetByIdEmpl(userId string, ctx context.Context) (restful.Employee, error)
	DeleteEmpl(userId string, ctx context.Context) error
	UpdateEmpl(newemployee restful.Employee, ctx context.Context) error
}
type Repository struct {
	EmplWork
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		EmplWork: NewEmplWorkPostgres(db),
	}
}
