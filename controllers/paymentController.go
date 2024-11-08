package controllers

import (
	"net/http"

	"github.com/gesangwidigdo/store-management/initializers"
	"github.com/gesangwidigdo/store-management/models"
	"github.com/gesangwidigdo/store-management/utils"
	"github.com/gin-gonic/gin"
)

type paymentInput struct {
	Transaction_id uint   `json:"transaction_id" binding:"required,numeric"`
	Payment_method string `json:"payment_method" binding:"required"`
}

func CreatePayment(c *gin.Context) {
	var paymentData paymentInput
	var transaction models.Transaction

	// bind data
	if bindErr := utils.BindData(&paymentData, c); !bindErr {
		return
	}

	// check transaction
	if selectedTransaction := initializers.DB.First(&transaction, paymentData.Transaction_id); selectedTransaction.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "transaction not found", "error", selectedTransaction.Error.Error(), c)
		return
	}

	if transaction.Status {
		utils.ReturnResponse(http.StatusBadRequest, "payment failed", "error", "transaction already paid", c)
		return
	}

	payment := models.Payment{
		Transaction_id: paymentData.Transaction_id,
		Payment_method: paymentData.Payment_method,
	}

	if result := initializers.DB.Create(&payment); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed create", "error", result.Error.Error(), c)
		return
	}

	// Update transaction status
	if result := initializers.DB.Model(&transaction).Where("id = ?", paymentData.Transaction_id).Update("status", true); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed update transaction status", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusCreated, "ok", "data", payment, c)
}
