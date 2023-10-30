package dummy

import (
	"context"
	"errors"
	"middle-developer-test/common/constants"
	"middle-developer-test/database"
	"middle-developer-test/dto"
	"middle-developer-test/model"
	"time"
)

type employeeRepo struct {
}

func NewEmployeeRepo() database.Employee {
	return &employeeRepo{}
}

func (r *employeeRepo) GetAllEmployeeData(ctx context.Context) ([]model.EmployeeData, error) {
	return []model.EmployeeData{
		{
			Id:        1,
			FirstName: "Data",
			LastName:  "1",
			Email:     "data1@email.com",
			HireDate:  time.Now().AddDate(0, -3, 0),
		},
		{
			Id:        2,
			FirstName: "Data",
			LastName:  "2",
			Email:     "data2@email.com",
			HireDate:  time.Now().AddDate(0, -2, 0),
		},
	}, nil
}

func (r *employeeRepo) GetEmployeeDataById(ctx context.Context, id int) (model.EmployeeData, error) {
	if id == 0 {
		return model.EmployeeData{}, errors.New(constants.NO_SQL_RESULT)
	} else {
		return model.EmployeeData{
			Id:        id,
			FirstName: "Data",
			LastName:  "Test",
			Email:     "datatest@email.com",
			HireDate:  time.Now().AddDate(0, -3, 0),
		}, nil
	}

}

func (r *employeeRepo) InsertEmployeeData(ctx context.Context, request dto.Employee) error {
	if request.FirstName == "Error" {
		return errors.New("failed insert")
	} else {
		return nil
	}
}

func (r *employeeRepo) UpdateEmployeeData(ctx context.Context, request dto.Employee) error {
	if request.Id == 0 {
		return errors.New("failed update")
	} else {
		return nil
	}
}

func (r *employeeRepo) DeleteEmployeeData(ctx context.Context, id int) error {
	if id == 0 {
		return errors.New("failed delete")
	} else {
		return nil
	}
}
