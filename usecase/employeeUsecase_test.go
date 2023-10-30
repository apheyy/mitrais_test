package usecase

import (
	"context"
	dummyDatabase "middle-developer-test/dummy/database"
	"middle-developer-test/model"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllEmployeeSuccess(t *testing.T) {
	uc := NewEmployeeUsecase(dummyDatabase.NewEmployeeRepo())
	respCode := uc.GetAllEmployee(context.Background())
	assert.Equal(t, respCode.Status, 200)
}

func TestGetEmployeeByIdSuccess(t *testing.T) {
	uc := NewEmployeeUsecase(dummyDatabase.NewEmployeeRepo())
	respCode := uc.GetEmployeeById(context.Background(), 1)
	assert.Equal(t, respCode.Status, http.StatusOK)
}

func TestGetEmployeeByIdFailed(t *testing.T) {
	uc := NewEmployeeUsecase(dummyDatabase.NewEmployeeRepo())
	respCode := uc.GetEmployeeById(context.Background(), 0)
	assert.Equal(t, respCode.Status, http.StatusInternalServerError)
	assert.NotEqual(t, respCode.Error, nil)
}

func TestInsertNewEmployeeSuccess(t *testing.T) {
	uc := NewEmployeeUsecase(dummyDatabase.NewEmployeeRepo())
	respCode := uc.InsertNewEmployee(
		context.Background(),
		model.UpsertEmployeeDataRequest{
			FirstName: "Testing",
			LastName:  "Data",
			Email:     "testingData@email.com",
			HireDate:  "2023-10-30",
		},
	)
	assert.Equal(t, respCode.Status, http.StatusOK)
}

func TestInsertNewEmployeeFailed(t *testing.T) {
	uc := NewEmployeeUsecase(dummyDatabase.NewEmployeeRepo())
	respCode := uc.InsertNewEmployee(
		context.Background(),
		model.UpsertEmployeeDataRequest{
			FirstName: "Error",
			LastName:  "Data",
			Email:     "testingData@email.com",
			HireDate:  "2023-10-30",
		},
	)
	assert.Equal(t, respCode.Status, http.StatusInternalServerError)
	assert.NotEqual(t, respCode.Error, nil)
}

func TestUpdateEmployeeByIdSuccess(t *testing.T) {
	uc := NewEmployeeUsecase(dummyDatabase.NewEmployeeRepo())
	respCode := uc.UpdateEmployeeById(
		context.Background(),
		model.UpsertEmployeeDataRequest{
			FirstName: "Testing",
			LastName:  "Data",
			Email:     "testingData@email.com",
			HireDate:  "2023-10-30",
		},
		1,
	)
	assert.Equal(t, respCode.Status, http.StatusOK)
}

func TestUpdateEmployeeByIdFailed(t *testing.T) {
	uc := NewEmployeeUsecase(dummyDatabase.NewEmployeeRepo())
	respCode := uc.UpdateEmployeeById(
		context.Background(),
		model.UpsertEmployeeDataRequest{
			FirstName: "Error",
			LastName:  "Data",
			Email:     "testingData@email.com",
			HireDate:  "2023-10-30",
		},
		0,
	)
	assert.Equal(t, respCode.Status, http.StatusBadRequest)
	assert.NotEqual(t, respCode.Error, nil)
}

func TestDeleteEmployeeByIdSuccess(t *testing.T) {
	uc := NewEmployeeUsecase(dummyDatabase.NewEmployeeRepo())
	respCode := uc.DeleteEmployeeById(
		context.Background(),
		1,
	)
	assert.Equal(t, respCode.Status, http.StatusOK)
}

func TestDeleteEmployeeByIdFailed(t *testing.T) {
	uc := NewEmployeeUsecase(dummyDatabase.NewEmployeeRepo())
	respCode := uc.DeleteEmployeeById(
		context.Background(),
		0,
	)
	assert.Equal(t, respCode.Status, http.StatusBadRequest)
	assert.NotEqual(t, respCode.Error, nil)
}
