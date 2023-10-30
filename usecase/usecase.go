package usecase

import (
	"context"
	"middle-developer-test/database"
	"middle-developer-test/model"
)

type (
	employeeUsecase struct {
		employeeRP database.Employee
	}

	Employee interface {
		GetAllEmployee(ctx context.Context) *model.ApiResponse
		GetEmployeeById(ctx context.Context, employeeId int) *model.ApiResponse
		InsertNewEmployee(ctx context.Context, request model.UpsertEmployeeDataRequest) *model.ApiResponse
		UpdateEmployeeById(ctx context.Context, request model.UpsertEmployeeDataRequest, employeeId int) *model.ApiResponse
		DeleteEmployeeById(ctx context.Context, employeeId int) *model.ApiResponse
	}
)
