package database

import (
	"context"
	"errors"
	"log"
	"middle-developer-test/common/constants"
	"middle-developer-test/dto"
	"middle-developer-test/model"

	"xorm.io/xorm"
)

func NewEmployeeRepo(db *xorm.Engine) Employee {
	return &employeeRepo{
		db: db,
	}
}

func (r *employeeRepo) GetAllEmployeeData(ctx context.Context) ([]model.EmployeeData, error) {
	resp := []model.EmployeeData{}
	listData := []dto.Employee{}
	sess := r.db.NewSession()
	defer sess.Close()
	err := sess.Find(&listData)
	if err != nil {
		log.Printf("[GetAllEmployeeData] fail to get data from employee table with error : %v", err)
	}

	if len(listData) < 1 {
		err = errors.New(constants.NO_SQL_RESULT)
		log.Printf("[GetAllEmployeeData] fail to get data from employee table with error : %v", err)
	} else {
		for _, data := range listData {
			resp = append(resp, MapEmployeeDtoIntoData(data))
		}
	}

	return resp, err
}

func (r *employeeRepo) GetEmployeeDataById(ctx context.Context, id int) (model.EmployeeData, error) {
	resp := model.EmployeeData{}
	data := dto.Employee{}
	sess := r.db.NewSession()
	defer sess.Close()
	has, err := sess.Where("id = ?", id).Get(&data)
	if err != nil {
		log.Printf("[GetEmployeeDataById] fail to get data from employee with id : %d, error : %v", id, err)
	}

	if !has {
		err = errors.New(constants.NO_SQL_RESULT)
		log.Printf("[GetEmployeeDataById] no data from employee with id : %d, error : %v", id, err)
	} else {
		log.Printf("[GetEmployeeDataById] success get data from employee with id : %d", id)
		resp = MapEmployeeDtoIntoData(data)
	}

	return resp, err
}

func (r *employeeRepo) InsertEmployeeData(ctx context.Context, request dto.Employee) error {
	sess := r.db.NewSession()
	defer sess.Close()

	_, err := sess.Insert(&request)
	if err != nil {
		return errors.New("[InsertEmployeeData] failed insert data into employee table")
	}

	err = sess.Commit()

	return err
}

func (r *employeeRepo) UpdateEmployeeData(ctx context.Context, request dto.Employee) error {
	sess := r.db.NewSession()
	defer sess.Close()

	_, err := sess.Where("id = ?", request.Id).Update(&request)
	if err != nil {
		return errors.New("failed update employee data")
	}

	err = sess.Commit()

	return err
}

func (r *employeeRepo) DeleteEmployeeData(ctx context.Context, id int) error {
	sess := r.db.NewSession()
	defer sess.Close()

	data := dto.Employee{}
	ids, err := sess.Where("id = ?", id).Delete(&data)
	if err != nil {
		return errors.New("failed delete employee data")
	}

	if ids == 0 {
		return errors.New(constants.NO_SQL_RESULT)
	}

	return err
}

func MapEmployeeDtoIntoData(data dto.Employee) model.EmployeeData {
	return model.EmployeeData{
		Id:        data.Id,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		HireDate:  data.HireDate,
	}
}
