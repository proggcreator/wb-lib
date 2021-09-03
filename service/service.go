package service

import (
	"context"

	restful "github.com/proggcreator/wb-lib/"
	repository "github.com/proggcreator/wb-lib/repository"
)

type EmplWork interface {
	CreateEmpl(employee restful.Employee, ctx context.Context) (string, error)
	GetAllEmpl(ctx context.Context) ([]restful.Employee, error)
	GetByIdEmpl(userId string, ctx context.Context) (restful.Employee, error)
	DeleteEmpl(userId string, ctx context.Context) error
	UpdateEmpl(newemployee restful.Employee, ctx context.Context) error
}
type Service struct {
	EmplWork
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		EmplWork: NewEmplWorkService(repos.EmplWork)}
}
