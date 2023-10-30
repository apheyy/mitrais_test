package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"middle-developer-test/common/constants"
	"middle-developer-test/config"
	"middle-developer-test/model"
	"middle-developer-test/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type employeeAPI struct {
	employeeUC usecase.Employee
}

func NewEmployeeAPI(employeeUC usecase.Employee) *employeeAPI {
	return &employeeAPI{
		employeeUC: employeeUC,
	}
}

func (employeeAPI *employeeAPI) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	funcName := "GetAllEmployee"
	log.Printf("[%s] Start get all employee from DB", funcName)

	resp := employeeAPI.employeeUC.GetAllEmployee(ctx)

	config.ResponseJSON(w, resp)
}

func (employeeAPI *employeeAPI) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	funcName := "GetEmployeeById"
	log.Printf("[%s] Start get employee data by id", funcName)

	path := mux.Vars(r)
	employeeId, _ := strconv.Atoi(path["id"])

	resp := employeeAPI.employeeUC.GetEmployeeById(ctx, employeeId)

	config.ResponseJSON(w, resp)
}

func (employeeAPI *employeeAPI) InsertNewEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	resp := &model.ApiResponse{}
	funcName := "InsertNewEmployee"
	log.Printf("[%s] Start insert employee data", funcName)

	body, _ := ioutil.ReadAll(r.Body)
	request := model.UpsertEmployeeDataRequest{}
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("[%s] Failed parse request with error : %+v ", funcName, err)

		resp = &model.ApiResponse{
			Status: http.StatusBadRequest,
			Error: &model.ErrorData{
				Message: constants.INVALID_PARAMETER_MESSAGE,
				Reason:  constants.INVALID_PARAMETER_REASON,
				Action:  constants.INVALID_PARAMETER_ACTION,
			},
		}
	}

	resp = employeeAPI.employeeUC.InsertNewEmployee(ctx, request)

	config.ResponseJSON(w, resp)
}

func (employeeAPI *employeeAPI) UpdateEmployeeById(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	resp := &model.ApiResponse{}
	funcName := "UpdateEmployeeById"
	log.Printf("[%s] Start update employee data", funcName)

	body, _ := ioutil.ReadAll(r.Body)
	request := model.UpsertEmployeeDataRequest{}

	path := mux.Vars(r)
	employeeId, _ := strconv.Atoi(path["id"])

	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("[%s] Failed parse request with error : %+v ", funcName, err)

		resp = &model.ApiResponse{
			Status: http.StatusInternalServerError,
			Error: &model.ErrorData{
				Message: constants.INVALID_PARAMETER_MESSAGE,
				Reason:  constants.INVALID_PARAMETER_REASON,
				Action:  constants.INVALID_PARAMETER_ACTION,
			},
		}
	}

	resp = employeeAPI.employeeUC.UpdateEmployeeById(ctx, request, employeeId)

	config.ResponseJSON(w, resp)
}

func (employeeAPI *employeeAPI) DeleteEmployeeById(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	funcName := "DeleteEmployeeById"
	log.Printf("[%s] Start delete employee data", funcName)

	path := mux.Vars(r)
	employeeId, _ := strconv.Atoi(path["id"])

	resp := employeeAPI.employeeUC.DeleteEmployeeById(ctx, employeeId)

	config.ResponseJSON(w, resp)
}
