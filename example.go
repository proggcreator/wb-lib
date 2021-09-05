package restful

import "database/sql"

func RetEmployee() Employee {
	return Employee{
		Id: sql.NullString{
			String: "57",
			Valid:  true},
		Name: sql.NullString{
			String: "John",
			Valid:  true},
		Last_name: sql.NullString{
			String: "Smith",
			Valid:  true},
		Patronymic: sql.NullString{
			String: "-",
			Valid:  true},
		Phone: sql.NullString{
			String: "89554857",
			Valid:  true},
		Position: sql.NullString{
			String: "manager",
			Valid:  true},
		Good_job_count: sql.NullInt64{
			Int64: 7,
			Valid: true},
	}
}
