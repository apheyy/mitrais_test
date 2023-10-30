package controller

import (
	"log"
	"middle-developer-test/common/constants"
	"middle-developer-test/config"
	"middle-developer-test/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func HandleAPI(cfg model.AppConfig) {
	muxRouter := mux.NewRouter().StrictSlash(true)
	r := muxRouter.PathPrefix(constants.MITRAIS).Subrouter()

	ng := negroni.New()
	ng.Use(negroni.HandlerFunc(config.CaptureNegroniHandler))

	uc := InitUsecase(cfg)
	apiModule := NewAPI(uc.employeeUC)

	r.HandleFunc(constants.EMPLOYEES, apiModule.employeeAPI.GetAllEmployee).Methods("GET")
	r.HandleFunc(constants.EMPLOYEES+constants.EMPLOYEE_ID, apiModule.employeeAPI.GetEmployeeById).Methods("GET")
	r.HandleFunc(constants.EMPLOYEES, apiModule.employeeAPI.InsertNewEmployee).Methods("POST")
	r.HandleFunc(constants.EMPLOYEES+constants.EMPLOYEE_ID, apiModule.employeeAPI.UpdateEmployeeById).Methods("PUT")
	r.HandleFunc(constants.EMPLOYEES+constants.EMPLOYEE_ID, apiModule.employeeAPI.DeleteEmployeeById).Methods("DELETE")

	log.Println("Loading API routes ....")
	ng.UseHandler(muxRouter)

	apiPort := 8000
	log.Println("api running on port:", apiPort)
	http.ListenAndServe(":"+strconv.Itoa(apiPort), ng)
}
