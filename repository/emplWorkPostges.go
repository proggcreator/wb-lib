package repository

import (
	"context"
	"database/sql"
	"fmt"

	restful "git.wildberries.ru/finance/general-documentation/go-intership-program/day-06"
	wbsql "git.wildberries.ru/finance/go-infrastructure/database/v2"
)

type EmplWorkPostgres struct {
	db *sql.DB
}

func NewEmplWorkPostgres(db *sql.DB) *EmplWorkPostgres {
	return &EmplWorkPostgres{db: db}
}

func (s *EmplWorkPostgres) CreateEmpl(empl restful.Employee, ctx context.Context) (string, error) {
	query := fmt.Sprintf("SELECT * FROM employees.employee_add(S1,$2,$3,$4,$5,$6,$7);")
	//wbsql.ExecQuery
	err := wbsql.ExecQuery(s.db, query, empl.Id, empl.Name, empl.Last_name, empl.Patronymic, empl.Phone, empl.Position, empl.Good_job_count)
	if err != nil {
		return "", err
	}
	return empl.Id, nil
}

func (s *EmplWorkPostgres) GetAllEmpl(ctx context.Context) ([]restful.Employee, error) {
	var lists []restful.Employee

	query := fmt.Sprintf("SELECT * FROM employees.get_all();")
	rows, err := wbsql.GetRows(s.db, query)
	if err != nil {
		//error
	}
	scanner := wbsql.InitScanner()
	for rows.Next() {
		var value1 restful.Employee

		scanner.ScanRows(rows).
			ScanSqlNullString("_id", &value1.Id).
			ScanSqlNullString("_name", &value1.Name).
			ScanSqlNullString("_last_name", &value1.Last_name).
			ScanSqlNullString("_patronymic", &value1.Patronymic).
			ScanSqlNullString("_phone", &value1.Phone).
			ScanSqlNullString("_position", &value1.Position).
			ScanSqlNullInt64("_good_job_count", &value1.Good_job_count)

		if scanner.Error() != nil {
			//error
		}
		lists = append(lists, value1)
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
