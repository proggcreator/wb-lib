package service

import (
	"context"

	restful "github.com/proggcreator/wb-Restful"
	"github.com/proggcreator/wb-Restful/repository"
)

type EmplWorkService struct {
	repo repository.EmplWork
}

func NewEmplWorkService(repo repository.EmplWork) *EmplWorkService {
	return &EmplWorkService{repo: repo}
}

func (s *EmplWorkService) CreateEmpl(empl restful.Employee, ctx context.Context) (string, error) {
	return s.repo.CreateEmpl(empl, ctx)
}

func (s *EmplWorkService) GetAllEmpl(ctx context.Context) ([]restful.Employee, error) {
	return s.repo.GetAllEmpl(ctx)
}
func (s *EmplWorkService) GetByIdEmpl(userId string, ctx context.Context) (restful.Employee, error) {
	return s.repo.GetByIdEmpl(userId, ctx)
}
func (s *EmplWorkService) DeleteEmpl(userId string, ctx context.Context) error {
	return s.repo.DeleteEmpl(userId, ctx)
}
func (s *EmplWorkService) UpdateEmpl(newemployee restful.Employee, ctx context.Context) error {
	return s.repo.UpdateEmpl(newemployee, ctx)
}
