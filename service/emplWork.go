package service

import (
	restful "github.com/proggcreator/wb-lib"
	repository "github.com/proggcreator/wb-lib/repository"
)

type EmplWorkService struct {
	repo repository.EmplWork
}

func NewEmplWorkService(repo repository.EmplWork) *EmplWorkService {
	return &EmplWorkService{repo: repo}
}

func (s *EmplWorkService) CreateEmpl(empl restful.Employee) (string, error) {
	return s.repo.CreateEmpl(empl)
}

func (s *EmplWorkService) GetAllEmpl() ([]restful.Employee, error) {
	return s.repo.GetAllEmpl()
}
func (s *EmplWorkService) GetByIdEmpl(userId string) (restful.Employee, error) {
	return s.repo.GetByIdEmpl(userId)
}
func (s *EmplWorkService) DeleteEmpl(userId string) error {
	return s.repo.DeleteEmpl(userId)
}
func (s *EmplWorkService) UpdateEmpl(newemployee restful.Employee) error {
	return s.repo.UpdateEmpl(newemployee)
}
