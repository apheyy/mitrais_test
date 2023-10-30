package database

import (
	"context"
	"middle-developer-test/dto"
	"middle-developer-test/model"

	"xorm.io/xorm"
)

type (
	employeeRepo struct {
		db *xorm.Engine
	}

	Employee interface {
		GetAllEmployeeData(ctx context.Context) ([]model.EmployeeData, error)
		GetEmployeeDataById(ctx context.Context, id int) (model.EmployeeData, error)
		InsertEmployeeData(ctx context.Context, request dto.Employee) error
		UpdateEmployeeData(ctx context.Context, request dto.Employee) error
		DeleteEmployeeData(ctx context.Context, id int) error
	}
)
