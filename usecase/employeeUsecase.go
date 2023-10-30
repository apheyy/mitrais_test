package usecase

import (
	"context"
	"log"
	"middle-developer-test/common/constants"
	"middle-developer-test/database"
	"middle-developer-test/dto"
	"middle-developer-test/model"
	"net/http"
	"time"
)

func NewEmployeeUsecase(employeeRP database.Employee) Employee {
	return &employeeUsecase{
		employeeRP: employeeRP,
	}
}

func (employeeUsecase *employeeUsecase) GetAllEmployee(ctx context.Context) *model.ApiResponse {
	funcName := "GetAllEmployee"
	data, err := employeeUsecase.employeeRP.GetAllEmployeeData(ctx)
	if err != nil {
		log.Printf("[%s] Failed find data on DB with error : %+v ", funcName, err)
		if err.Error() == constants.NO_SQL_RESULT {
			return &model.ApiResponse{
				Status: http.StatusNotFound,
				Error: &model.ErrorData{
					Message: constants.NO_DATA_FOUND_MESSAGE,
					Reason:  constants.NO_DATA_FOUND_REASON,
					Action:  constants.NO_DATA_FOUND_ACTION,
				},
			}
		} else {
			return &model.ApiResponse{
				Status: http.StatusInternalServerError,
				Error: &model.ErrorData{
					Message: constants.INTERNAL_SERVER_ERROR_MESSAGE,
					Reason:  constants.INTERNAL_SERVER_ERROR_REASON,
					Action:  constants.INTERNAL_SERVER_ERROR_ACTION,
				},
			}
		}
	}

	return &model.ApiResponse{
		Status: http.StatusOK,
		Data:   data,
	}
}

func (employeeUsecase *employeeUsecase) GetEmployeeById(ctx context.Context, employeeId int) *model.ApiResponse {
	funcName := "GetEmployeeById"
	data, err := employeeUsecase.employeeRP.GetEmployeeDataById(ctx, employeeId)
	if err != nil {
		log.Printf("[%s] Failed find data on DB with error : %+v ", funcName, err)

		return &model.ApiResponse{
			Status: http.StatusInternalServerError,
			Error: &model.ErrorData{
				Message: constants.ID_NOT_FOUND_MESSAGE,
				Reason:  constants.ID_NOT_FOUND_REASON,
				Action:  constants.ID_NOT_FOUND_ACTION,
			},
		}
	}

	return &model.ApiResponse{
		Status: http.StatusOK,
		Data:   data,
	}
}

func (employeeUsecase *employeeUsecase) InsertNewEmployee(ctx context.Context, request model.UpsertEmployeeDataRequest) *model.ApiResponse {
	funcName := "InsertNewEmployee"

	hireDate, err := time.Parse(constants.HIRE_DATE_PARSE, request.HireDate)
	if err != nil {
		return &model.ApiResponse{
			Status: http.StatusBadRequest,
			Error: &model.ErrorData{
				Message: constants.INVALID_PARAMETER_MESSAGE,
				Reason:  constants.INVALID_PARAMETER_REASON,
				Action:  constants.INVALID_PARAMETER_ACTION,
			},
		}
	}

	err = employeeUsecase.employeeRP.InsertEmployeeData(ctx, dto.Employee{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		HireDate:  hireDate,
	})
	if err != nil {
		log.Printf("[%s] Failed insert new data to DB with error : %+v ", funcName, err)

		return &model.ApiResponse{
			Status: http.StatusInternalServerError,
			Error: &model.ErrorData{
				Message: constants.FAILED_INSERT_DATA_MESSAGE,
				Reason:  constants.FAILED_INSERT_DATA_REASON,
				Action:  constants.FAILED_INSERT_DATA_ACTION,
			},
		}
	}

	return &model.ApiResponse{
		Status: http.StatusOK,
		Data:   constants.SUCCESS_INSERT_EMPLOYEE_DATA,
	}
}

func (employeeUsecase *employeeUsecase) UpdateEmployeeById(ctx context.Context, request model.UpsertEmployeeDataRequest, employeeId int) *model.ApiResponse {
	funcName := "UpdateEmployeeById"

	hireDate, err := time.Parse(constants.HIRE_DATE_PARSE, request.HireDate)
	if err != nil {
		return &model.ApiResponse{
			Status: http.StatusBadRequest,
			Error: &model.ErrorData{
				Message: constants.INVALID_PARAMETER_MESSAGE,
				Reason:  constants.INVALID_PARAMETER_REASON,
				Action:  constants.INVALID_PARAMETER_ACTION,
			},
		}
	}

	existingData, err := employeeUsecase.employeeRP.GetEmployeeDataById(ctx, employeeId)
	if err != nil {
		log.Printf("[%s] Failed find data on DB with error : %+v ", funcName, err)

		return &model.ApiResponse{
			Status: http.StatusBadRequest,
			Error: &model.ErrorData{
				Message: constants.ID_NOT_FOUND_MESSAGE,
				Reason:  constants.ID_NOT_FOUND_REASON,
				Action:  constants.ID_NOT_FOUND_ACTION,
			},
		}
	}

	isRequestChangeSameAsExistingData := CompareExistingDataWithRequest(existingData, request, hireDate)

	if isRequestChangeSameAsExistingData {
		return &model.ApiResponse{
			Status: http.StatusBadRequest,
			Error: &model.ErrorData{
				Message: constants.UPDATE_DATA_ALREADY_EXISTS_MESSAGE,
				Reason:  constants.UPDATE_DATA_ALREADY_EXISTS_REASON,
				Action:  constants.UPDATE_DATA_ALREADY_EXISTS_ACTION,
			},
		}
	}

	err = employeeUsecase.employeeRP.UpdateEmployeeData(ctx, dto.Employee{
		Id:        employeeId,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		HireDate:  hireDate,
	})
	if err != nil {
		log.Printf("[%s] Failed update data to DB with error : %+v ", funcName, err)

		return &model.ApiResponse{
			Status: http.StatusInternalServerError,
			Error: &model.ErrorData{
				Message: constants.FAILED_INSERT_DATA_MESSAGE,
				Reason:  constants.FAILED_INSERT_DATA_REASON,
				Action:  constants.FAILED_INSERT_DATA_ACTION,
			},
		}
	}

	return &model.ApiResponse{
		Status: http.StatusOK,
		Data:   constants.SUCCESS_UPDATE_EMPLOYEE_DATA,
	}
}

func (employeeUsecase *employeeUsecase) DeleteEmployeeById(ctx context.Context, employeeId int) *model.ApiResponse {
	funcName := "DeleteEmployeeById"

	data, err := employeeUsecase.employeeRP.GetEmployeeDataById(ctx, employeeId)
	if err != nil {
		log.Printf("[%s] Failed find data on DB with error : %+v ", funcName, err)

		return &model.ApiResponse{
			Status: http.StatusBadRequest,
			Error: &model.ErrorData{
				Message: constants.ID_NOT_FOUND_MESSAGE,
				Reason:  constants.ID_NOT_FOUND_REASON,
				Action:  constants.ID_NOT_FOUND_ACTION,
			},
		}
	}

	err = employeeUsecase.employeeRP.DeleteEmployeeData(ctx, data.Id)
	if err != nil {
		log.Printf("[%s] Failed delete employee data with error : %+v ", funcName, err)

		return &model.ApiResponse{
			Status: http.StatusInternalServerError,
			Error: &model.ErrorData{
				Message: constants.FAILED_DELETE_DATA_MESSAGE,
				Reason:  constants.FAILED_DELETE_DATA_REASON,
				Action:  constants.FAILED_DELETE_DATA_ACTION,
			},
		}

	}

	return &model.ApiResponse{
		Status: http.StatusOK,
		Data:   constants.SUCCESS_DELETE_EMPLOYEE_DATA,
	}
}

func CompareExistingDataWithRequest(oldData model.EmployeeData, newData model.UpsertEmployeeDataRequest, hireDate time.Time) bool {
	return oldData.FirstName == newData.FirstName && oldData.LastName == newData.LastName && oldData.Email == newData.Email && oldData.HireDate.Equal(hireDate)
}
