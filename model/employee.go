package model

import "time"

type EmployeeData struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	HireDate  time.Time `json:"hire_date"`
}

type UpsertEmployeeDataRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	HireDate  string `json:"hireDate"`
}
