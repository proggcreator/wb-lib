package restful

func RetEmployee() Employee {
	return Employee{
		Id:             "57",
		Name:           "John",
		Last_name:      "Smith",
		Patronymic:     "-",
		Phone:          "89554857",
		Position:       "manager",
		Good_job_count: 7,
	}
}
