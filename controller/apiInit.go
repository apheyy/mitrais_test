package controller

import (
	"middle-developer-test/config"
	"middle-developer-test/database"
	"middle-developer-test/model"
	"middle-developer-test/usecase"
)

func NewAPI(
	employeeUC usecase.Employee,
) *APIModule {
	return &APIModule{
		employeeAPI: NewEmployeeAPI(employeeUC),
	}
}

type Usecase struct {
	employeeUC usecase.Employee
}

func InitUsecase(cfg model.AppConfig) (uc Usecase) {
	db := config.InitDB(cfg)

	employeeRP := database.NewEmployeeRepo(db)

	uc.employeeUC = usecase.NewEmployeeUsecase(employeeRP)

	return uc
}
