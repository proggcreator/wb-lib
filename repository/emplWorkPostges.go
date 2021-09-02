package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	restful "github.com/proggcreator/wb-Restful"
)

type EmplWorkPostgres struct {
	db *sqlx.DB
}

func NewEmplWorkPostgres(db *sqlx.DB) *EmplWorkPostgres {
	return &EmplWorkPostgres{db: db}
}

func (s *EmplWorkPostgres) CreateEmpl(empl restful.Employee, ctx context.Context) (string, error) {
	query := fmt.Sprintf("SELECT * FROM employees.employee_add(S1,$2,$3,$4,$5,$6,$7);")
	_, err := s.db.Exec(query, empl.Id, empl.Name, empl.Last_name, empl.Patronymic, empl.Phone, empl.Position, empl.Good_job_count)
	if err != nil {
		return "", err
	}
	return empl.Id, nil
}

func (s *EmplWorkPostgres) GetAllEmpl(ctx context.Context) ([]restful.Employee, error) {
	var lists []restful.Employee
	query := fmt.Sprintf("SELECT * FROM employees.get_all();")
	err := s.db.Select(&lists, query)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (s *EmplWorkPostgres) GetByIdEmpl(userId string, ctx context.Context) (restful.Employee, error) {

	var empl restful.Employee
	query := fmt.Sprintf("SELECT * FROM employees.get_all($1);")
	err := s.db.Get(&empl, query, userId)
	if err != nil {
		return restful.Employee{}, err
	}
	return empl, nil
}

func (s *EmplWorkPostgres) DeleteEmpl(userId string, ctx context.Context) error {
	query := fmt.Sprintf("SELECT * FROM employees.employee_remove($1);")
	_, err := s.db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmplWorkPostgres) UpdateEmpl(newemployee restful.Employee, ctx context.Context) error {
	query := fmt.Sprintf("SELECT * FROM employees.employee_remove($1,$2,$3,$4,$5,$6,$7);")
	_, err := s.db.Exec(query, newemployee.Id, newemployee.Name, newemployee.Last_name,
		newemployee.Patronymic, newemployee.Phone, newemployee.Position, newemployee.Good_job_count)
	if err != nil {
		return err
	}

	return nil
}
