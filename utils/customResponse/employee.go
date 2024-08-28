package customresponse

import (
	"time"

	"github.com/gesangwidigdo/store-management/models"
	// "gorm.io/gorm"
)

type EmployeeResponse struct {
	ID               uint      `json:"id"`
	Employee_name    string    `json:"employee_name"`
	Gender           string    `json:"gender"`
	Telephone_number string    `json:"telephone_number"`
	Username         string    `json:"username"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	// DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type ForeignEmployeeResponse struct {
	ID            uint   `json:"id"`
	Employee_name string `json:"employee_name"`
	Gender        string `json:"gender"`
	Username      string `json:"username"`
}

func EmployeeResponseData(employee models.Employee) EmployeeResponse {
	return EmployeeResponse{
		ID:               employee.ID,
		Employee_name:    employee.Employee_name,
		Gender:           employee.Gender,
		Telephone_number: employee.Telephone_number,
		Username:         employee.Username,
		CreatedAt:        employee.CreatedAt,
		UpdatedAt:        employee.UpdatedAt,
		// DeletedAt:        employee.DeletedAt,
	}
}

func ForeignEmployeeResponseData(employee models.Employee) ForeignEmployeeResponse {
	return ForeignEmployeeResponse{
		ID:            employee.ID,
		Employee_name: employee.Employee_name,
		Gender:        employee.Gender,
		Username:      employee.Username,
	}
}
