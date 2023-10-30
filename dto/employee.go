package dto

import "time"

type Employee struct {
	Id        int       `xorm:"pk autoincr 'id'"`
	FirstName string    `xorm:"'first_name'"`
	LastName  string    `xorm:"'last_name'"`
	Email     string    `xorm:"'email'"`
	HireDate  time.Time `xorm:"TIMESTAMPZ 'hire_date'"`
}
