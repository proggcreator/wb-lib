package service

import (
	restful "github.com/proggcreator/wb-lib/"
	repository "github.com/proggcreator/wb-lib/repository"
)

type EmplWork interface {
	CreateEmpl(employee restful.Employee) (string, error)
	GetAllEmpl() ([]restful.Employee, error)
	GetByIdEmpl(userId string) (restful.Employee, error)
	DeleteEmpl(userId string) error
	UpdateEmpl(newemployee restful.Employee) error
}
type Service struct {
	EmplWork
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		EmplWork: NewEmplWorkService(repos.EmplWork)}
}
