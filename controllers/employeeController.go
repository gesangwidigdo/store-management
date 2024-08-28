package controllers

import (
	"net/http"

	"github.com/gesangwidigdo/store-management/initializers"
	"github.com/gesangwidigdo/store-management/models"
	"github.com/gesangwidigdo/store-management/utils"
	customresponse "github.com/gesangwidigdo/store-management/utils/customResponse"
	"github.com/gin-gonic/gin"
)

type employeeInput struct {
	Employee_name    string `json:"employee_name" binding:"required,alpha"`
	Gender           string `json:"gender" binding:"required,alpha"`
	Telephone_number string `json:"telephone_number" binding:"required,min=10"`
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
}

func CreateEmployee(c *gin.Context) {
	var employeeData employeeInput

	// bind data
	if err := utils.BindData(&employeeData, c); !err {
		return
	}

	// hash pwd
	hashedPwd, err := utils.HashPassword(employeeData.Password)
	if err != nil {
		utils.ReturnResponse(http.StatusBadRequest, "hash failed", "error", err.Error(), c)
		return
	}

	employee := models.Employee{
		Employee_name:    employeeData.Employee_name,
		Gender:           employeeData.Gender,
		Telephone_number: employeeData.Telephone_number,
		Username:         employeeData.Username,
		Password:         hashedPwd,
	}

	if result := initializers.DB.Create(&employee); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed create", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusCreated, "ok", "data", employee, c)
}

func GetAllEmployee(c *gin.Context) {
	var employees []models.Employee

	if result := initializers.DB.Find(&employees); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed get data", "error", result.Error.Error(), c)
		return
	}

	var employeeResponses []customresponse.EmployeeResponse
	for _, emp := range employees {
		employeeResponses = append(employeeResponses, customresponse.EmployeeResponseData(emp))
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", employeeResponses, c)
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	if result := initializers.DB.First(&employee, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed get data", "error", result.Error.Error(), c)
		return
	}

	employeeResponse := customresponse.EmployeeResponseData(employee)

	utils.ReturnResponse(http.StatusOK, "ok", "data", employeeResponse, c)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employeeData employeeInput

	if err := utils.BindData(&employeeData, c); !err {
		return
	}

	var employee models.Employee
	if result := initializers.DB.First(&employee, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "error", result.Error.Error(), c)
		return
	}

	// hash pwd
	hashedPwd, err := utils.HashPassword(employeeData.Password)
	if err != nil {
		utils.ReturnResponse(http.StatusBadRequest, "hash failed", "error", err.Error(), c)
		return
	}

	// new data
	newData := models.Employee{
		Employee_name:    employeeData.Employee_name,
		Gender:           employeeData.Gender,
		Telephone_number: employeeData.Telephone_number,
		Username:         employeeData.Username,
		Password:         hashedPwd,
	}

	if result := initializers.DB.Model(&employee).Updates(newData); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update data", "error", result.Error.Error(), c)
		return
	}

	var updatedEmployee models.Employee
	if result := initializers.DB.First(&updatedEmployee, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update data", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "new_data", updatedEmployee, c)
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee

	if result := initializers.DB.Delete(&employee, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed delete data", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "", nil, c)
}
