package controllers

import (
	"net/http"

	"github.com/gesangwidigdo/store-management/initializers"
	"github.com/gesangwidigdo/store-management/models"
	"github.com/gesangwidigdo/store-management/utils"
	"github.com/gin-gonic/gin"
)

type ProductTransactionInput struct {
	Transaction_id uint    `json:"transaction_id" binding:"required,numeric"`
	Product_id     uint    `json:"product_id" binding:"required,numeric"`
	Quantity       int     `json:"quantity" binding:"required,numeric"`
}

func CreateProductTransaction(c *gin.Context) {
	var PTData ProductTransactionInput

	// bind data
	if err := utils.BindData(&PTData, c); !err {
		return
	}

	// Get Product Data
	var product models.Product
	if err := initializers.DB.First(&product, PTData.Product_id).Error; err != nil {
		utils.ReturnResponse(http.StatusBadRequest, "retrieve product data failed", "error", "data not found", c)
		return
	}

	// Check if product's stock decent or not
	if product.Stock < PTData.Quantity {
		utils.ReturnResponse(http.StatusBadRequest, "failed add data", "error", "not enough stock", c)
		return
	}

	// get transaction data
	var transaction models.Transaction
	if err := initializers.DB.First(&transaction, PTData.Transaction_id); err.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "retrieve transaction data failed", "error", "data not found", c)
		return
	}

	// calculate total
	totalPrice := product.Price * float64(PTData.Quantity)
	productTransaction := models.ProductTransaction{
		Transaction_id: PTData.Transaction_id,
		Product_id: PTData.Product_id,
		Quantity: PTData.Quantity,
		Total: totalPrice,
	}

	if result := initializers.DB.Create(&productTransaction); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed create", "error", result.Error.Error(), c)
		return
	}

	// decrease product's stock
	if err := initializers.DB.Model(&models.Product{}).Where("id = ?", PTData.Product_id).Update("stock", product.Stock - PTData.Quantity); err.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update data", "error", err.Error.Error(), c)
		return
	}

	// update grand total
	if err := initializers.DB.Model(&models.Transaction{}).Where("id = ?", PTData.Transaction_id).Update("grand_total", transaction.Grand_total + totalPrice); err.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update transaction data", "error", err.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusCreated, "ok", "data", productTransaction, c)
}