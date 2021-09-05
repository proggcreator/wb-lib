package restful

import (
	"database/sql"
)

type Employee struct {
	Id             sql.NullString
	Name           sql.NullString
	Last_name      sql.NullString
	Patronymic     sql.NullString
	Phone          sql.NullString
	Position       sql.NullString
	Good_job_count sql.NullInt64
}
