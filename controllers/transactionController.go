package controllers

import (
	"net/http"

	"github.com/gesangwidigdo/store-management/initializers"
	"github.com/gesangwidigdo/store-management/models"
	"github.com/gesangwidigdo/store-management/utils"
	customresponse "github.com/gesangwidigdo/store-management/utils/customResponse"
	"github.com/gin-gonic/gin"
)

type transactionInput struct {
	Employee_id uint `json:"employee_id" binding:"required,numeric"`
}

func CreateTransaction(c *gin.Context) {
	var transactionData transactionInput

	// bind data
	if err := utils.BindData(&transactionData, c); !err {
		return
	}

	transaction := models.Transaction{
		Employee_id: transactionData.Employee_id,
	}

	if result := initializers.DB.Create(&transaction); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, utils.CREATE_FAILED, "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusCreated, "ok", "data", transaction, c)
}

func GetAllTransaction(c *gin.Context) {
	var transactions []models.Transaction

	if result := initializers.DB.Preload("Employee").Find(&transactions); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, utils.GET_FAILED, "error", result.Error.Error(), c)
		return
	}

	var transactionResponses []customresponse.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, customresponse.ToTransactionResponse(transaction))
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", transactionResponses, c)
}

func GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	if result := initializers.DB.Preload("Employee").First(&transaction, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, utils.GET_FAILED, "error", result.Error.Error(), c)
		return
	}

	transactionResponse := customresponse.ToTransactionResponse(transaction)

	utils.ReturnResponse(http.StatusOK, "ok", "data", transactionResponse, c)
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	if result := initializers.DB.Delete(&transaction, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, utils.GET_FAILED, "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "", nil, c)
}