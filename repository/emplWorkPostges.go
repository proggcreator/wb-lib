package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	wbsql "git.wildberries.ru/finance/go-infrastructure/database/v2"
	wbsqlreq "git.wildberries.ru/finance/go-infrastructure/sql-requests/v2"
	restful "github.com/proggcreator/wb-lib"
)

type EmplWorkPostgres struct {
	db wbsql.DbConnecter
}

func NewEmplWorkPostgres(db wbsql.DbConnecter) *EmplWorkPostgres {
	return &EmplWorkPostgres{db: db}
}

//transformation of interface
func retPgDb(s *EmplWorkPostgres) (*sql.DB, error) {
	pgdb, err := s.db.Connection()
	if err != nil {
		return nil, err
	}
	sqlDb, ok := pgdb.(*sql.DB)
	if !ok {
		return nil, err
	}
	return sqlDb, nil
}

func (s *EmplWorkPostgres) CreateEmpl(empl restful.Employee) (string, error) {
	query := fmt.Sprintf("SELECT * FROM employees.employee_add(S1,$2,$3,$4,$5,$6,$7);")

	//get *sql.DB
	sqlDb, err := retPgDb(s)
	if err != nil {
		return "", err
	}
	err = wbsql.ExecQuery(sqlDb, query, empl.Id.String, empl.Name.String, empl.Last_name.String,
		empl.Patronymic.String, empl.Phone.String, empl.Position.String, empl.Good_job_count.Int64)
	if err != nil {
		return "", err
	}
	return empl.Id.String, nil
}

func (s *EmplWorkPostgres) GetAllEmpl() ([]restful.Employee, error) {
	var lists []restful.Employee

	//get *sql.DB
	sqlDb, err := retPgDb(s)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM employees.get_all();")
	rows, err := wbsql.GetRows(sqlDb, query)
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

func (s *EmplWorkPostgres) GetByIdEmpl(userId string) (restful.Employee, error) {
	//query := fmt.Sprintf("SELECT * FROM employees.get_all($1);")

	//get *sql.DB
	sqlDb, err := retPgDb(s)
	if err != nil {
		return restful.Employee{}, err
	}
	resultQuery := wbsqlreq.PgSqlBuilder().
		SetDbName(s.db.Name()).
		SetSelect().ForProcedure("employees.get_all").
		AddParameter("_id", userId).Build()

	rows, err := wbsql.GetRows(sqlDb, resultQuery)
	if err != nil {
		//error
	}
	scanner := wbsql.InitScanner()
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
		return restful.Employee{}, err
	}
	return value1, nil

}

func (s *EmplWorkPostgres) DeleteEmpl(userId string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	executor := wbsql.CreateSqlExecuter(s.db)
	query := fmt.Sprintf("SELECT * FROM employees.employee_remove($1);")
	err := executor.ExecuteContext(ctx, query, userId)

	if err != nil {
		return err
	}
	return nil
}

func (s *EmplWorkPostgres) UpdateEmpl(newemployee restful.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	executor := wbsql.CreateSqlExecuter(s.db)
	query := fmt.Sprintf("SELECT * FROM employees.employee_remove($1,$2,$3,$4,$5,$6,$7);")
	err := executor.ExecuteContext(ctx, query, newemployee.Id.String, newemployee.Name.String,
		newemployee.Last_name.String, newemployee.Patronymic.String, newemployee.Phone.String,
		newemployee.Position.String, newemployee.Good_job_count.Int64)
	if err != nil {
		return err
	}

	return nil
}
